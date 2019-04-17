import { helper } from '@ember/component/helper';

export function formatHosts([hosts]/*, hash*/) {
  return hosts.replace(/,/g, '====>');
}

export default helper(formatHosts);
