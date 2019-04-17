import { helper } from '@ember/component/helper';

export function toLower([value]/*, hash*/) {
    let str = value || '';
    return str.toLowerCase();
}

export default helper(toLower);
