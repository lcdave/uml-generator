<template>
  <svg
    @mousemove="onMouseOver"
    class="sparkline"
    :width="width"
    :height="height"
    :stroke-width="stroke"
  >
    <line
      :x1="this.lineX"
      y1="0"
      :x2="this.lineX"
      :y2="this.height"
      style="stroke: rgb(255, 0, 0); stroke-width: 2"
    />
    <path class="sparkline--line" :d="shape" fill="none"></path>
    <path
      class="sparkline--fill"
      :d="[shape, fillEndPath].join(' ')"
      stroke="none"
    ></path>
  </svg>
</template>

<script>
export default {
  props: ["data"],
  data() {
    return {
      stroke: 2,
      width: 256,
      height: 30,
      hover: false,
      coordinates: [],
      lineX: 0,
    };
  },
  computed: {
    shape() {
      const stroke = this.stroke;
      const width = this.width;
      const height = this.height - stroke * 2;
      const data = this.data || [];
      const highestPoint = Math.max.apply(null, data);
      const totalPoints = this.data.length - 1;
      data.forEach((respTime, n) => {
        const x = (n / totalPoints) * width + stroke;
        const y = height - (respTime / highestPoint) * height + stroke;
        this.coordinates.push({ x, y, respTime });
      });

      if (!this.coordinates[0]) {
        return (
          "M 0 " +
          this.stroke +
          " L 0 " +
          this.stroke +
          " L " +
          this.width +
          " " +
          this.stroke
        );
      }

      const path = [];
      this.coordinates.forEach((point) =>
        path.push(["L", point.x, point.y].join(" "))
      );
      return ["M" + this.coordinates[0].x, this.coordinates[0].y, ...path].join(
        " "
      );
    },
    fillEndPath() {
      return `V ${this.height} L 4 ${this.height} Z`;
    },
  },
  methods: {
    closest(coors, num) {
      var i = 0;
      var minDiff = 1000;
      var coor;
      for (i in coors) {
        var m = Math.abs(num - coors[i].x);
        if (m < minDiff) {
          minDiff = m;
          coor = coors[i];
        }
      }
      return coor;
    },
    onMouseOver(e) {
      const rect = e.target.getBoundingClientRect();
      let x = e.clientX - rect.left; //x position within the element.
      const width = this.width;

      // clamp x
      x = Math.max(0, Math.min(x, width));

      let closest = this.closest(this.coordinates, x);

      console.log(closest);

      this.lineX = closest.x;

      //line.visible = true;
      //line.x = closest.x;
    },
    //onMouseLeave(e)
  },
};
</script>

<style>
svg {
  stroke: #9dcbe5;
  fill: rgba(31, 140, 235, 0.06);
  transition: all 1s ease-in-out;
}
svg path {
  box-sizing: border-box;
}
</style>
