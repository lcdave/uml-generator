<template>
  <div class="quickstat">
    <div class="quickstat__icon">
      <slot name="icon" />
    </div>
    <div class="quickstat__content">
      <div class="quickstat__value">
        <slot name="value" />
      </div>
      <div class="quickstat__description">
        <slot name="description" />
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" src="../scss/components/_quickstat.scss"></style>
