<template>
  <div>
    <div class="container">
      <!-- <SpWelcome /> -->
      <h2> {{ networkName }}</h2>
      <SpTokenSend :address="address" :refresh="true" />
      <SpTransferList :address="address" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Index',
  data(){
    return {
      networkName: process.env.VUE_APP_CHAIN_NAME

    }
  },
  computed: {
    address() {
      return this.$store.getters['common/wallet/address']
    },
  },
}
</script>
