import Service from '@ember/service';
import {getOwner} from '@ember/application';

export default Service.extend({
    currentUser: null,
    loadCurrentUser() {
        let appInstance = getOwner(this);
        let store = appInstance.lookup('service:store');
        return store.request.post('/v1/auth').then((data)=>{
            if(!data || !data.uid){
                throw new Error('登录错误，请联系管理员');
            }
            this.set('currentUser', data);
        }).catch((e)=>{
            throw new Error(e);
        });
    },
    setToken({token, token_expire}){
        this.set('currentUser.token', token);
        this.set('currentUser.token_expire', token_expire);
    }
});
