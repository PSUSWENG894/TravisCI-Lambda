<template>
    <div>
        <p>This is a simple component for capturing build data.</p>
    </div>
</template>

<script>
export default {
  props: {
    buildNumber: String,
    buildUsername: String,
    buildTimeout: Number,
    buildRuntimeVersions: String,
    buildTestPhase: String,
    buildCompilePhase: String,
    buildDeployPhase: String,
    buildWasPreviousDeploy: String,
    buildWasBuildCanceled: String
  }
};
</script>

<style scoped>
</style>