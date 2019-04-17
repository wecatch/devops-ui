import { helper } from '@ember/component/helper';

export function deployStatus([status]/*, hash*/) {
  let msg = "";
  switch (status){
    case 'new':
      msg = '准备部署';
      break;
    case 'doing':
      msg = '正在部署';
      break;
    case 'fail':
      msg = '部署失败';
      break;
    case 'success':
      msg = '部署成功';
      break;
    default:
      msg = '未知状态'
      break;
  }

  return msg;
}

export default helper(deployStatus);
