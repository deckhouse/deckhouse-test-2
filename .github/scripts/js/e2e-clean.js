const {abortFailedE2eCommand} = require("./constants");

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
        const branch = needsContext[key].outputs['ref'] || '';

        if (!branch || !stateDir || !ranFor || !connectStr || !artifactName || !runId || !startCommentId) {
          core.warn(`Incorrect outputs for ${key}: ${JSON.stringify(outputs)}`)
        }

        const splitRunFor = ranFor.replace(';', ' ');

        return `
<!--- failed_clusters_start ${ranFor} -->
E2e for ${splitRunFor} was failed. Use:
  \`ssh -i ~/.ssh/e2e-id-rsa ${connectStr}\` - connect for debugging;

  \`${abortFailedE2eCommand} ${branch} ${ranFor} ${runId} ${artifactName} ${stateDir} ${startCommentId}\` - for abort failed cluster
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
