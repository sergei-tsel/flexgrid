<template>
  <div class="grid-container">
    <Header class="grid-element-head"/>
    <slot></slot>
    <Footer class="grid-element-foot"/>
  </div>
</template>

<script setup lang="ts">

import Header from "~/components/Header.vue";
</script>

<style scoped>
.grid-container {
  display: grid;
  width: 100%;
  height: 100vh;
  grid-template-areas:
        "H H H   H H H   H H H   H H H   H H H   H H H"
        "H H H   H H H   H H H   H H H   H H H   H H H"
        "I I I   I I I   I I I   I I I   I I I   I I I"

        "I I I   I I I   I I I   I I I   I I I   I I I"
        "I I I   I I I   I I I   I I I   I I I   I I I"
        "E E E   E E B   B B B   C C C   C D D   D D D"

        "E E E   E E B   B B B   C C C   C D D   D D D"
        "E E E   E E A   A A A   A A A   A D D   D D D"
        "E E E   E E A   A A A   A A A   A D D   D D D"

        "E E E   E E A   A A A   A A A   A D D   D D D"
        "E E E   E E A   A A A   A A A   A D D   D D D"
        "E E E   E E A   A A A   A A A   A D D   D D D"

        "E E E   E E A   A A A   A A A   A D D   D D D"
        "G G G   G G G   G G G   G G G   G G G   G G G"
        "G G G   G G G   G G G   G G G   G G G   G G G"

        "G G G   G G G   G G G   G G G   G G G   G G G"
        "F F F   F F F   F F F   F F F   F F F   F F F"
        "F F F   F F F   F F F   F F F   F F F   F F F";
  grid-template-rows:    2fr 0.25fr 0.75fr   1fr 0.5fr 1.5fr   0.75fr 1.25fr 1fr   1fr 1.25fr 0.75fr   1.5fr 0.5fr 1fr   0.75fr 0.25fr 2fr;
  grid-template-columns: 2fr 0.25fr 0.75fr   1fr 0.5fr 1.5fr   0.75fr 1.25fr 1fr   1fr 1.25fr 0.75fr   1.5fr 0.5fr 1fr   0.75fr 0.25fr 2fr;
}

.grid-element-auth {
  grid-area: A;
}

.grid-element-button {
  grid-area: B;
}

.grid-element-control {
  grid-area: C;
}

.grid-element-daily {
  grid-area: D;
}

.grid-element-easy {
  grid-area: E;
}

.grid-element-foot {
  grid-area: F;
}

.grid-element-ground {
  grid-area: G;
}

.grid-element-head {
  grid-area: H;
}

.grid-element-intro {
  grid-area: I;
}
</style>
