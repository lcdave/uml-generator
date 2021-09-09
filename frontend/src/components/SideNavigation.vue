<template>
  <div class="side-navigation">
    <div class="side-navigation__logo">
      <img src="../assets/img/logo.svg" alt="Atticode logo" />
    </div>
    <div class="side-navigation__cta">
      <ul>
        <li>
          <router-link to="/">
            <span class="icon-home var_grey is-active" />
          </router-link>
        </li>
        <li>
          <router-link to="/service/all/"
            ><span class="icon-bar-chart var_grey"
          /></router-link>
        </li>
        <li>
          <router-link to="/logs/all/"
            ><span class="icon-history var_grey"
          /></router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped src="../scss/components/_side-navigation.scss" />
