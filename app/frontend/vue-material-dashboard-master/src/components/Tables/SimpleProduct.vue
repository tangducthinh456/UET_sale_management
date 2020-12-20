<template>
  <div>
  
    <md-table v-model="products" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }" @click="showDialog = true; setForm(item)">
        <md-table-cell md-label="Product ID">{{ item.product_id }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.product_name }}</md-table-cell>
        <md-table-cell md-label="Price ($)">{{ item.price }}</md-table-cell>
        <md-table-cell md-label="Quantity">{{ item.quantity }}</md-table-cell>
      </md-table-row>
    </md-table>

    <div>
      <md-dialog-alert
      :md-active.sync="successAlert"
      md-content="Success!"
      md-confirm-text="Ok" />

    <md-dialog-alert
      :md-active.sync="failAlert"
      md-title="Error!"
      md-content="Fail!" />
    <md-dialog :md-active.sync="showDialog">
      
      <md-dialog-title>Product</md-dialog-title>
      
      <md-dialog-content>
      
      <div style="margin: 30px">
      <form class="md-layout">
          
      <md-field>
          <label>Product Name</label>
         <md-input v-model="productNameForm"></md-input>   
     </md-field>
     
     <md-autocomplete   v-model="selectedGroup" :md-options="createdGroup">
      <label>Group</label>
    </md-autocomplete>
     
    <md-field >
          <label>Cost ($)</label>
         <md-input v-model="costForm"></md-input>
     </md-field>

    <md-field >
          <label>Price ($)</label>
         <md-input v-model="priceForm"></md-input>
     </md-field>

    <md-field >
          <label>Quantity</label>
         <md-input v-model="quantityForm"></md-input>
     </md-field>

     <md-field >
          <label>Brand</label>
         <md-input v-model="brandForm"></md-input>
     </md-field>

     <md-field >
          <label>Position</label>
         <md-input v-model="positionForm"></md-input>
     </md-field>

     <md-field >
          <label>Description</label>
         <md-input v-model="descriptionForm"></md-input>
     </md-field>

     <md-field >
          <label>Created At</label>
         <md-input v-model="createdAtDisplay"></md-input>
     </md-field>

     <md-field >
          <label>Created User</label>
         <md-input v-model="createdUserForm"></md-input>
     </md-field>
     
     <md-dialog-actions>
        <md-dialog-alert
      :md-active.sync="successAlert"
      md-content="Your post has been deleted!"
      md-confirm-text="Cool!" />

    <md-dialog-alert
      :md-active.sync="failAlert"
      md-title="Post created!"
      md-content="Your post <strong>Material Design is awesome</strong> has been created." />
        <md-button class="md-primary" @click="showDialog = false;updateProduct()">Update</md-button>
        <md-button class="md-primary" @click="showDialog = false;deleteProduct()">Delete</md-button>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        
      </md-dialog-actions>
      </form>
      
      </div>
       </md-dialog-content>
      <!-- <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button  class="md-primary" @click="showDialog = false">Save</md-button>
      </md-dialog-actions> -->
    </md-dialog>

    
  </div>




  </div>
</template>

<script>
import axios from 'axios';

var group = [];
 var groupFull = [];

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
              prod.push(varib[i]);
              //prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/products?page_size=1000", true);
      xhttp.send(); 


      var getGroup = function() {
          var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText)
            while(group.length > 0){ 
              group.pop();
              groupFull.pop();
            }
            for (var i = 0; i < response.length; i++){
              group.push(response[i].group_name);
              groupFull.push(response[i]);
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/groups", true);
      xhttp.send(); 
      };
      getGroup();
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
      products: prod,
      createdGroup: group,
      showDialog: false,

      failAlert: false,
      successAlert: false,

      productIDForm: null,
      productNameForm: null,
      selectedGroup: null,
      groupIdForm: null,
      costForm: null,
      priceForm: null,
      quantityForm: null,
      brandForm: null,
      positionForm: null,
      descriptionForm: null,
      createdAtForm: null,
      createdUserForm: null,
      creatdUserIDForm: null,
      createdAtDisplay: null   
    };
  },
  props: ['productName', 'createdId', 'selectedDateFrom', 'selectedDateTo', 'groupId'],
  watch: {
  productName: 'reload',
  },
  methods: {
    setForm(item){
      this.productIDForm = item.product_id;
      this.productNameForm = item.product_name;
      this.selectedGroup = item.group.group_name;
      this.groupIdForm = item.group.group_id;
      this.costForm = item.cost.toString();
      this.priceForm = item.price.toString();
      this.quantityForm = item.quantity;
      this.brandForm = item.brand;
      this.positionForm = item.position;
      this.descriptionForm = item.description;
      this.createdAtDisplay = item.created_at.replace(/T/g,' ').replace(/\+(.*)$/g,'');
      this.createdAtForm = item.created_at;
      this.createdUserForm = item.created_user.name;
      this.createdUserIDForm = item.created_user_id;
    },
    updateProduct(){
      axios.put('http://localhost:8081/api/products/' + this.productIDForm,{
             product_name: this.productNameForm,
             group_id: this.groupIdForm,
             created_user_id: this.createdUserIDForm,
             cost: parseFloat(this.costForm),
             price: parseFloat(this.priceForm),
             quantity: parseInt(this.quantityForm),
             brand: this.brandForm,
             position: this.positionForm,
             description: this.descriptionForm,
             created_at: this.createdAtForm
          }).then((response) =>{

             if (response.status == 200){
               this.$set(this.$data, 'successAlert','true');
             }
             else {
               this.$set(this.$data, 'failAlert','true');
             }
          }).catch((error) =>{
            this.$set(this.$data, 'failAlert','true');
          })
          console.log(parseFloat(this.costForm) );    
    },
    deleteProduct(){
      axios.delete('http://localhost:8081/api/products/' + this.productIDForm)
      .then((response) =>{
             if (response.status == 200){
               this.$set(this.$data, 'successAlert','true');
             }
             else {
               this.$set(this.$data, 'failAlert','true');
             }
          }).catch((error) =>{
            this.$set(this.$data, 'failAlert','true');
          })
    },

    reload(){
      
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
              prod.push(varib[i]);
              //prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      var url = "http://localhost:8081/api/products?page_size=1000&filter=product_nameLIKE%25" + this.productName + "%25";
      if (this.groupId) {
           url += "%3Bgroup_id=" + this.groupId; 
      }
      if (this.createdId) {
           url += "%3Bcreated_id=" + this.createdId; 
      }
      if (this.selectedDateFrom) {
           console.log(this.selectedDateFrom)
           url += "%3Bcreated_at>=" + this.selectedDateFrom.toISOString();
      }
      if (this.selectedDateTo) {
           url += "%3Bcreated_at<=" + this.selectedDateTo.toISOString();
      }
      
      //xhttp.open("GET", "http://localhost:8081/api/products?page_size=1000&filter=product_nameLIKE%25" + this.productName + "%25%3Bgroup_id=" + this.groupId + "%3Bcreated_user_id=" + this.createdId + "%3Bcreated_at>=" + this.selectedDateFrom.toISOString() + "%3Bcreated_at<=" + this.selectedDateTo.toISOString(), true);
      xhttp.open("GET",url);
      xhttp.send(); 

    
    }
  }
};
</script>
