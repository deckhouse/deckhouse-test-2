const WORKFLOW_STATUS_RUNNING = 'in_progress';
const WORKFLOW_STATUS_COMPLETED = 'completed';
const MAX_ATTEMPTS = 60;
const TIMEOUT_BETWEEN_ATTEMPT = 1000 * 30;
const MAX_ITEMS_PER_PAGE = 100;

// Обертка для захвата зависимостей
module.exports = ({ github, context, core }) => {
  async function isWorkflowCompleted(branch, workflowName) {
    try {
      const { data } = await github.rest.actions.listWorkflowRunsForRepo({
        owner: context.repo.owner,
        repo: context.repo.repo,
        branch: branch,
        per_page: MAX_ITEMS_PER_PAGE,
      });

      const activeRuns = data.workflow_runs.filter(
        (run) => run.name === workflowName && run.status === WORKFLOW_STATUS_RUNNING
      );

      if (activeRuns.length > 0) {
        core.info(`🔄 Active '${workflowName}' workflows found, waiting...`);
        return false;
      }

      const completedRun = data.workflow_runs.find(
        (run) => run.name === workflowName && run.status === WORKFLOW_STATUS_COMPLETED
      );

      if (!completedRun) {
        core.setFailed('❌ No completed workflow found');
        return false;
      }

      return completedRun.conclusion === 'success';
    } catch (error) {
      core.setFailed(`🔥 Error: ${error.message}`);
      return false;
    }
  }

  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  return async function waitForWorkflowIsCompleted(
    branchName,
    workflowName,
    maxAttempts = MAX_ATTEMPTS,
    timeoutBetweenAttempt = TIMEOUT_BETWEEN_ATTEMPT
  ) {
    for (let i = 0; i < maxAttempts; i++) {
      core.info(`🚀 Attempt ${i + 1}/${maxAttempts}`);
      const isReady = await isWorkflowCompleted(branchName, workflowName);

      if (isReady) {
        core.info('✅ Workflow completed successfully!');
        return true;
      }

      await sleep(timeoutBetweenAttempt);
    }

    core.setFailed('⌛ Timeout waiting for workflow completion');
    throw new Error('Max attempts reached');
  };
};
