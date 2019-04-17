import { helper } from '@ember/component/helper';
import { htmlSafe } from '@ember/string';
import AnsiUp from 'ansi-up';

let ansi_up = new AnsiUp;

export function formatLogMessage([message]/*, hash*/) {
    if (message){
        return htmlSafe(ansi_up.ansi_to_html(message));
    }else {
        return message;
    }
}

export default helper(formatLogMessage);
