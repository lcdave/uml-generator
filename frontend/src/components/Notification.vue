<template>
  <transition name="fade" v-on:enter="enter">
    <div class="notification" :class="'notification' + variant" v-if="show">
      <div class="notification__message">{{ message }}</div>
    </div>
  </transition>
</template>

<script>
export default {
  data() {
    return {
      show: false,
    };
  },
  props: {
    variant: String,
    message: String,
  },
  methods: {
    fadeMe: function () {
      this.show = !this.show;
    },

    enter: function () {
      var that = this;

      setTimeout(function () {
        that.show = false;
      }, 3000); // hide the message after 3 seconds
    },
  },
};
</script>

<style lang="scss" src="../scss/components/_notification.scss" scoped></style>
