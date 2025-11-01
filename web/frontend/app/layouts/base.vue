<template>
  <div class="grid-container">
    <SiteHeader class="grid-element-h"/>
    <slot></slot>
    <SiteFooter class="grid-element-f"/>
  </div>
</template>

<script setup lang="ts">
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
        "E E E   E E B   B B B   C C C   C C D   D D D"

        "E E E   E E A   A A A   A A A   A A D   D D D"
        "E E E   E E A   A A A   A A A   A A D   D D D"
        "E E E   E E A   A A A   A A A   A A D   D D D"

        "E E E   E E A   A A A   A A A   A A D   D D D"
        "E E E   E E A   A A A   A A A   A A D   D D D"
        "E E E   E E A   A A A   A A A   A A D   D D D"

        "E E E   E E A   A A A   A A A   A A D   D D D"
        "G G G   G G G   G G G   G G G   G G G   G G G"
        "G G G   G G G   G G G   G G G   G G G   G G G"

        "G G G   G G G   G G G   G G G   G G G   G G G"
        "F F F   F F F   F F F   F F F   F F F   F F F"
        "F F F   F F F   F F F   F F F   F F F   F F F";
  grid-template-rows:    2fr 0.25fr 0.75fr   1fr 0.5fr 1.5fr   0.75fr 1.25fr 1fr   1fr 1.25fr 0.75fr   1.5fr 0.5fr 1fr   0.75fr 0.25fr 2fr;
  grid-template-columns: 2fr 0.25fr 0.75fr   1fr 0.5fr 1.5fr   0.75fr 1.25fr 1fr   1fr 1.25fr 0.75fr   1.5fr 0.5fr 1fr   0.75fr 0.25fr 2fr;
}

.grid-element-a {
  grid-area: A;
}

.grid-element-b {
  grid-area: B;
}

.grid-element-c {
  grid-area: C;
}

.grid-element-d {
  grid-area: D;
}

.grid-element-e {
  grid-area: E;
}

.grid-element-f {
  grid-area: F;
}

.grid-element-g {
  grid-area: G;
}

.grid-element-h {
  grid-area: H;
}

.grid-element-i {
  grid-area: I;
}
</style>
