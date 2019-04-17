import { helper } from '@ember/component/helper';

export function formatCommitId([commitId]/*, hash*/) {
    let length = commitId.length
    return commitId.substr(length - 6, length);
}

export default helper(formatCommitId);
