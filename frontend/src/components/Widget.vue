<template>
  <div
    class="widget"
    :class="{ ['var_' + size]: size, var_quickstat: quickStat }"
  >
    <div class="widget__header">
      <h3 class="widget__title">{{ title }}</h3>
      <router-link to="/service/new" class="iconlink" v-if="hasIcon">
        <span class="icon-plus var_green var_cta" />
      </router-link>
    </div>
    <div class="widget__content">
      <slot name="content" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    size: Number,
    quickStat: Boolean,
    hasIcon: Boolean,
  },
};
</script>

<style lang="scss" src="../scss/components/_widget.scss"></style>
