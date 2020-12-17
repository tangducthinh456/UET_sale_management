<template>
  <div class="content">
    <div class="md-layout">
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
        <md-card class="md-card-plain">
          <md-card-header data-background-color="green">
            <h4 class="title">Filter</h4>
            <p class="category">Here is a filter for this table</p>
          </md-card-header>
          <md-card-content>
            <simple-filter :productName.sync="productName"></simple-filter>
          </md-card-content>
        </md-card>
      </div>
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100">
        <md-card>
          <md-card-header data-background-color="blue">
            <h4 class="title">Product</h4>
            <p class="category">List product in shop</p>
          </md-card-header>
          <md-card-content>
            <simple-product table-header-color="blue" :product-name="productName"></simple-product>
          </md-card-content>
        </md-card>
      </div>
    </div>
  </div>
</template>

<script>
import { SimpleFilter, SimpleProduct } from "@/components";
//import Filter from '../components/Tables/Filter.vue';

export default {
  data: function () {
        return {
            productName: null,
            selectedGroup: "soap",
            createdUser: "manager_01",
            selectedDateFrom : null,
            selectedDateTo : null
       };
  },
  methods: {
    getFilter(pd){
      this.productName = pd;
      console.log("aaaaaaaaaaaaaaaaa");
    }
  },
  components: {
    SimpleFilter,
    SimpleProduct
  }
};
</script>
