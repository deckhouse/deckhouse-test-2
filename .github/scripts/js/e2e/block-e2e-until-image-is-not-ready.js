import core from '@actions/core';
import { Octokit } from '@octokit/rest';

const WORKFLOW_NAME = 'Sleep job';
const WORKFLOW_STATUS_RUNNING = 'in_progress';
const WORKFLOW_STATUS_COMPLETED = 'completed';
const MAX_ATTEMPTS = 60;
const TIMEOUT_BETWEEN_ATTEMPT = 1000 * 30; // 10 секунд
const MAX_ITEMS_PER_PAGE = 100;

const token = process.env.GITHUB_TOKEN;
const repo = process.env.REPO;
const prName = context.payload.inputs.pull_request_ref;
const octokit = new Octokit({ auth: token });

/**
 * @param {string} repo 
 * @param {string} branch 
 * @returns {Promise<boolean>}
 */
async function isReadyToE2E(repo, branch) {
  try {
    const [owner, repoName] = repo.split('/');
    const { data } = await octokit.actions.listWorkflowRunsForRepo({
      owner,
      repo: repoName,
      branch: branch,
      per_page: MAX_ITEMS_PER_PAGE,
    });

    // Проверяем наличие активных workflow 'Sleep job'
    const activeRuns = data.workflow_runs.filter(run => 
      run.name === WORKFLOW_NAME && run.status === WORKFLOW_STATUS_RUNNING
    );

    if (activeRuns.length > 0) {
      console.log(`Есть активные джобы '${WORKFLOW_NAME}', ждём их завершения.`);
      return false;
    }

    console.log(`Активных джоб '${WORKFLOW_NAME}' не найдено, проверяем завершение первой джобы.`);

    // Проверка первой завершенной джобы 'Sleep job'
    const completedRun = data.workflow_runs.find(run => 
      run.name === WORKFLOW_NAME && run.status === WORKFLOW_STATUS_COMPLETED
    );
    
    if (completedRun) {
      if (completedRun.conclusion === 'success') {
        console.log('Первая джоба закончилась успешно.');
        return true;
      } else {
        console.error('Первая джоба закончилась с ошибкой.');
        core.setFailed('Актуального образа нет, первая джоба завершилась с ошибкой.');
        return false;
      }
    } else {
      console.error('Первая завершенная джоба не найдена.');
      core.setFailed('Не найдена завершенная джоба для анализа.');
      return false;
    }
  } catch (error) {
    core.setFailed(error.message);
    return false;
  }
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

(async () => {
  for (let i = 0; i < MAX_ATTEMPTS; i++) {
    console.log(`Попытка ${i + 1} из ${MAX_ATTEMPTS}`);
    const isReady = await isReadyToE2E(repo, prName);
    if (isReady) {
      // добавить код для проверки образа, если джоба завершилась успешно.
      return;
    }
    await sleep(TIMEOUT_BETWEEN_ATTEMPT);
  }
  core.setFailed('Не удалось дождаться завершения джоб за допустимое количество попыток.');
})();