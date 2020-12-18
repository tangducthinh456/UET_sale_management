<template>
  <div>
    <md-table v-model="products" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="Product ID">{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Price ($)">{{ item.price }}</md-table-cell>
        <md-table-cell md-label="Quantity">{{ item.quantity }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
var prod = [];
export default {
  name: "simple-product",
  created: function(){
       var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            while(prod.length > 0){ 
              prod.pop();
            }
            //console.log("aaaaaaaaa");
            var response = JSON.parse(this.responseText)
            var varib = response.products
            for (var i = 0; i < varib.length; i++){
              prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/products?page_size=1000", true);
      xhttp.send(); 
  },
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      selected: [],
      products: prod
    };
  },
  props: ['productName', 'createdId', 'selectedDateFrom', 'selectedDateTo', 'groupId'],
  watch: {
  productName: function (newValue) {
    //console.log(this.selectedDateFrom.toISOString(), this.selectedDateTo.toISOString());
       var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            while(prod.length > 0){ 
              prod.pop();
            }
            //console.log("aaaaaaaaa");
            var response = JSON.parse(this.responseText)
            var varib = response.products
            for (var i = 0; i < varib.length; i++){
              prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/products?page_size=1000&filter=product_nameLIKE%25" + this.productName + "%25%3Bgroup_id=" + this.groupId + "%3Bcreated_user_id=" + this.createdId + "%3Bcreated_at>=" + this.selectedDateFrom.toISOString() + "%3Bcreated_at<=" + this.selectedDateTo.toISOString(), true);
      xhttp.send(); 
    }
    
  },
  
  

  
};
</script>
