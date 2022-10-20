const {abortFailedE2eCommand} = require("./constants");
const {checkAbortE2eCluster} = require("./e2e-clean");
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

  let slashCommand = dispatchPullRequestCommand(argv, core, context);
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

  return await startWorkflow({github, context, core,
    workflow_id: workflow.ID,
    ref: workflow.targetRef,
    inputs: {
      comment_id: '' + response.data.id,
      ...slashCommand.inputs
    },
  });
}

/**
 *
 * @param {string[]} argv - slash command arguments [0] arg is name of command
 * @param {object} core - github core object
 * @param {object} context - github core object
 * @return {object}
 */
function dispatchPullRequestCommand(argv, core, context){
  const command = argv[0];
  core.debug(`Command is ${argv[0]}`)
  core.debug(`argv is ${JSON.stringify(argv)}`)
  switch (command) {
    case abortFailedE2eCommand:
      return checkAbortE2eCluster(argv, context)
  }

  return null;
}

module.exports = {
  runSlashCommandForPullRequest,
}
