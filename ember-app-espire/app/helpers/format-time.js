import { helper } from '@ember/component/helper';
import moment from 'moment';

export function formatTime([time, timeEnd]/*, hash*/) {
  if (time && timeEnd){
    return moment(timeEnd).from(time);
  }
  return moment(time).format('YYYY-MM-DD hh:mm:ss');
}

export default helper(formatTime);
