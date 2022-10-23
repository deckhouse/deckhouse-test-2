const {abortFailedE2eCommand} = require("../constants");

/**
 * Build additional info about failed e2e test
 * Contains information about
 *
 * @param {object} jobs - GitHub needsContext context
 * @returns {string}
 */
function buildFailedE2eTestAdditionalInfo({ needsContext, core }){
  const connectStrings = Object.getOwnPropertyNames(needsContext).
  filter((k) => k.startsWith('run_')).
  map((key, _i, _a) => {
    const result = needsContext[key].result;
    if (result === 'failure' || result === 'cancelled') {
      if (needsContext[key].outputs){
        const outputs = needsContext[key].outputs;

        if(!outputs['failed_cluster_stayed']){
          return null;
        }

        // ci_commit_branch
        const connectStr = outputs['ssh_master_connection_string'] || '';
        const ranFor = outputs['ran_for'] || '';
        const runId = outputs['run_id'] || '';
        const startCommentId = outputs['start_e2e_comment_id'] || '';
        const artifactName = outputs['state_artifact_name'] || '';
        const stateDir = needsContext[key].outputs['state_dir'] || '';
        const ci_commit_ref_name = needsContext[key].outputs['ci_commit_ref_name'] || '';
        const pull_request_ref = needsContext[key].outputs['pull_request_ref'] || '';
        const pull_request_sha = needsContext[key].outputs['pull_request_sha'] || '';

        const argv = [
          abortFailedE2eCommand,
          ci_commit_ref_name,
          pull_request_ref,
          pull_request_sha,
          ranFor,
          runId,
          artifactName,
          stateDir,
          startCommentId,
        ]

        const shouldArgc = argv.length
        const argc = argv.filter(v => !!v).length

        if (shouldArgc !== argc) {
          core.debug(`Incorrect outputs for ${key} ${shouldArgc} != ${argc}: ${JSON.stringify(argv)}; ${JSON.stringify(outputs)}`)
        }

        const splitRunFor = ranFor.replace(';', ' ');

        return `
<!--- failed_clusters_start ${ranFor} -->
E2e for ${splitRunFor} was failed. Use:
  \`ssh -i ~/.ssh/e2e-id-rsa ${connectStr}\` - connect for debugging;

  \`${argv.join(' ')}\` - for abort failed cluster
<!--- failed_clusters_end ${ranFor} -->

`
      }
    }

    return null;
  }).filter((v) => !!v)

  if (connectStrings.length === 0) {
    return "";
  }

  return "\r\n" + connectStrings.join("\r\n") + "\r\n";
}

module.exports = {
  buildFailedE2eTestAdditionalInfo,
}
