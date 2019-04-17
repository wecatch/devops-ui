import { helper } from '@ember/component/helper';

export function extractProjectName([url]/*, hash*/) {
  if (url && typeof url === "string"){
    let array = url.split('/');
    if (array.length > 0) {
      return array[array.length - 1];
    }
  }
  return 'url';
}

export default helper(extractProjectName);
