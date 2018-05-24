import Vue from 'vue'
import BuildInformation from '@/components/BuildInformation'

describe('BuildInformation.vue', () => {
    
    it('should render correct contents', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm).toBeTruthy()
    });
    
    it('should Show the last build number', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm.buildNumber).toBe('A_BUILD_NUMBER')
    });
    
    it('should Show who initiated the build', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm.buildUsername).toBe('Username')
    });
    
    it('should Show build timeout', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm.buildTimeout).toBe(2000)
    });

    it('should list configuration runtime versions', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm.buildRuntimeVersions).toBe('2.2.080')
    });

    it('should list job phases', () => {
        const Constructor = Vue.extend(BuildInformation)
        const vm = new Constructor().$mount()
        expect(vm.buildTestPhase).toBe(true)
        expect(vm.buildCompilePhase).toBe(true)
        expect(vm.buildDeployPhase).toBe(true)
        expect(vm.buildWasPreviousDeploy).toBe(true)
        expect(vm.buildWasBuildCanceled).toBe(true)
    });
});
