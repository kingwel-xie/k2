import { sysHttp } from '/@/utils/http/axios';

export function getServerInfo() {
  return sysHttp.get<any>(
    { url: '/server-monitor' },
    { errorMessageMode: 'none', isTransformResponse: false },
  );
}
