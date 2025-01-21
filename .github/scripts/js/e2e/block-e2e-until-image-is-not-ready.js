const WORKFLOW_NAME = 'Build and test for dev branches';
const WORKFLOW_STATUS_RUNNING = 'in_progress';
const WORKFLOW_STATUS_COMPLETED = 'completed';
const MAX_ATTEMPTS = 60;
const TIMEOUT_BETWEEN_ATTEMPT = 1000 * 30; // 10 second
const MAX_ITEMS_PER_PAGE = 100;

/**
* @param {string} branch 
* @returns {Promise<boolean>}
*/
async function isReadyToE2E(branch) {
  try {
    const { data } = await github.rest.actions.listWorkflowRunsForRepo({
      owner: context.repo.owner,
      repo: context.repo.repo,
      branch: branch,
      per_page: MAX_ITEMS_PER_PAGE,
    });

    // Checking for active workflow 'Build and test for dev branches'
    const activeRuns = data.workflow_runs.filter(run => run.name === WORKFLOW_NAME && run.status === WORKFLOW_STATUS_RUNNING);
    if (activeRuns.length > 0) {
      console.log(`There are active '${WORKFLOW_NAME}' jobs, wait for them to complete.`);
      return false;
    }

    // Checking the status of the first task 'Build and test for dev branches'
    console.log(`No active jobs '${WORKFLOW_NAME}' were found, checking status first job.`);
    const completedRun = data.workflow_runs.find(run => run.name === WORKFLOW_NAME && run.status === WORKFLOW_STATUS_COMPLETED);
    
    if (completedRun) {
      if (completedRun.conclusion === 'success') {
        console.log('The first job was completed successfully.');
        return true;
      } else {
        console.error('The first job ended with an error.');
        core.setFailed('There is no current image; the first job finished with an error.');
        return false;
      }
    } else {
      core.setFailed('Job not found');
      return false;
    }
  } catch (error) {
    core.setFailed(error.message);
    return false;
  }
}

const branchName = context.payload.inputs.ci_commit_ref_name;
const prNum = context.payload.inputs.issue_number;
console.log(`Run check for branch: ${branchName} PR: ${context.payload.repository.html_url}/pull/${prNum}`);

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function isWorkflowReadyToE2E() {
  for (let i = 0; i < MAX_ATTEMPTS; i++) {
    console.log(`Attempt number ${i + 1} of ${MAX_ATTEMPTS}`);
    const isReady = await isReadyToE2E(branchName);
    if (isReady) {
      return;
    }
    await sleep(TIMEOUT_BETWEEN_ATTEMPT);
  }
  core.setFailed('Failed to wait for the job to complete within the allowed number of attempts.');
};

module.exports = {
    isWorkflowReadyToE2E
}