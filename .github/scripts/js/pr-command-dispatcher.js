const {abortFailedE2eCommand} = require("./constants");
const {commentCommandRecognition} = require("./comments");
const {extractCommandFromComment, reactToComment, startWorkflow} = require("./ci");

/**
 * Use pull request comment to determine a workflow to run.
 *
 * @param {object} inputs
 * @param {object} inputs.github - A pre-authenticated octokit/rest.js client with pagination plugins.
 * @param {object} inputs.context - An object containing the context of the workflow run.
 * @param {object} inputs.core - A reference to the '@actions/core' package.
 * @returns {Promise<void|*>}
 */
async function runSlashCommandForPullRequest({ github, context, core }) {
  const event = context.payload;
  const comment_id = event.comment.id;
  core.debug(`Event: ${JSON.stringify(event)}`);

  const arg = extractCommandFromComment(event.comment.body)
  if(arg.err) {
    return core.info(`Ignore comment: ${arg.err}.`);
  }

  const { argv } = arg;

  let slashCommand = dispatchPullRequestCommand(argv);
  if (!slashCommand) {
    return core.info(`Ignore comment: command ${argv[0]} not found.`);
  }

  if (slashCommand.err) {
    return core.setFailed(`Cannot start workflow: ${slashCommand.err}`);
  }

  core.info(`Command detected: ${JSON.stringify(slashCommand)}`);

  const { workflow } = slashCommand;
  // Git ref is malformed.
  if (!workflow.targetRef) {
    core.setFailed('targetRef is missed');
    return await reactToComment({github, context, comment_id, content: 'confused'});
  }

  // Git ref is malformed.
  if (!workflow.ID) {
    core.setFailed('workflowID is missed');
    return await reactToComment({github, context, comment_id, content: 'confused'});
  }

  core.info(`Use ref '${workflow.targetRef}' for workflow.`);

  // React with rocket emoji!
  await reactToComment({github, context, comment_id, content: 'rocket'});

  // Add new issue comment and start the requested workflow.
  core.info('Add issue comment to report workflow status.');
  let response = await github.rest.issues.createComment({
    owner: context.repo.owner,
    repo: context.repo.repo,
    issue_number: event.issue.number,
    body: commentCommandRecognition(event.comment.user.login, argv[0])
  });

  if (response.status !== 201) {
    return core.setFailed(`Cannot start workflow: ${JSON.stringify(response)}`);
  }

  const commentInfo = {
    issue_id: '' + event.issue.id,
    issue_number: '' + event.issue.number,
    comment_id: '' + response.data.id,
  };

  return await startWorkflow({github, context, core,
    workflow_id: workflow.ID,
    ref: workflow.targetRef,
    inputs: {
      ...commentInfo,
      ...slashCommand.inputs
    },
  });
}

/**
 *
 * @param {string[]} argv - slash command arguments [0] arg is name of command
 * @return {object}
 */
function dispatchPullRequestCommand(argv){
  const command = argv[0];
  switch (command) {
    case abortFailedE2eCommand:
      return checkAbortE2eCluster(argv)
  }

  return null;
}

/**
 *
 * @param {string[]} parts - slash command strings
 * @return {object}
 */
function checkAbortE2eCluster(parts){
  const command = parts[0];
  if (command !== abortFailedE2eCommand) {
    return null;
  }

  if (parts.length !== 6) {
    let err = 'clean failed e2e cluster should have 6 arguments'
    switch (parts.length){
      case 5:
        err = 'state dir is required';
        break;
      case 4:
        err = 'artifact name and state dir are required';
        break;
      case 3:
        err = 'run id, artifact name and state dir are required';
        break;
      case 2:
        err = 'ran for (provider, layout, cri, k8s version), run id, artifact name, and state dir are required';
        break;
      case 1:
        err = 'full_ref, ran for (provider, layout, cri, k8s version), run id, artifact name, and state dir are required';
        break;
    }
    return {err};
  }

  const ranForSplit = parts[2].split(',').map(v => v.trim()).filter(v => !!v);
  if (ranForSplit.length !== 4) {
    let err = '"ran for" argument should have 4 parts';
    switch (ranForSplit.length) {
      case 3:
        err = 'k8s version is required';
        break;
      case 2:
        err = 'cri and k8s version are required';
        break;
      case 1:
        err = 'layout, cri and k8s version are required';
        break;
      case 0:
        err = 'provider, layout, cri and k8s version are required';
        break;
    }

    return {err};
  }

  const provider = ranForSplit[0];

  return {
    isDestroyFailedE2e: true,
    workflow: {
      ID: `e2e-clean-${provider}.yml`,
      targetRef: parts[1],
    },
    inputs: {
      run_id: parts[3],
      state_artifact_name: parts[4],
      state_dir: parts[5],

      layout: ranForSplit[1],
      cri: ranForSplit[2],
      k8s_version: ranForSplit[3],
    },
  };
}


module.exports = {
  runSlashCommandForPullRequest,
}
