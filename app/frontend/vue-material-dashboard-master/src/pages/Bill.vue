<template>
  <div class="content">
    <div class="md-layout">
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
     
      <div>
      <md-dialog-alert
      :md-active.sync="successAlert"
      md-content="Bill has been created!"
      md-confirm-text="Ok" />

    <md-dialog-alert
      :md-active.sync="failAlert"
      md-title="Error!"
      md-content="Created Fail!" />
    <md-dialog :md-active.sync="showDialog">
      
      <md-dialog-title>Customer</md-dialog-title>
      
      <md-dialog-content>
      
      <div style="margin: 30px">
      <form class="md-layout">
     
     <md-autocomplete   v-model="selectedCustomer" :md-options="createdCustomer">
      <label>Customer</label>
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
     
     <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button type="submit" class="md-primary" @click="showDialog = false;formSubmit()">Save</md-button>
      </md-dialog-actions>
      </form>
      
      </div>
       </md-dialog-content>
      <!-- <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button  class="md-primary" @click="showDialog = false">Save</md-button>
      </md-dialog-actions> -->
    </md-dialog>

    <md-button class="md-primary md-raised" @click="showDialog = true"><md-icon>library_add</md-icon>Create Bill</md-button>
  </div>
      
        
      
        <md-card class="md-card-plain">
          <md-card-header data-background-color="green">
            <h4 class="title">Filter</h4>
            <p class="category">Here is a filter for this table</p>
          </md-card-header>
          <md-card-content>
            <bill-filter  :createdId.sync="createdId" :customerId.sync="customerId" :selectedDateFrom.sync="selectedDateFrom" :selectedDateTo.sync="selectedDateTo"></bill-filter>
          </md-card-content>
        </md-card>
      </div>
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100">
        <md-card>
          <md-card-header data-background-color="blue">
            <h4 class="title">Bill</h4>
            <p class="category">List bill</p>
          </md-card-header>
          
          <md-card-content>
            <simple-bill table-header-color="blue" :created-id="createdId" :customer-id="customerId" :selected-date-from="selectedDateFrom" :selected-date-to="selectedDateTo"></simple-bill>
          </md-card-content>
        </md-card>
      </div>
    </div>
  </div>
</template>

<script >

import axios from 'axios';

var customer = [];
 var customerFull = [];
  var user = [];
  var userFull = [];

import { BillFilter, SimpleBill } from "@/components";
//import Filter from '../components/Tables/Filter.vue';

export default {
  created: function(){
      var getCustomer = function() {
          var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText).customers
            while(customer.length > 0){ 
              customer.pop();
              customerFull.pop();
            }
            for (var i = 0; i < response.length; i++){
              customer.push(response[i].customer_name);
              customerFull.push(response[i]);
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/customers?page_size=100", true);
      xhttp.send(); 
      };
      getCustomer();

      var getUser = function() {
          var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText)
            //console.log(this.responseText)
            for (var i = 0; i < response.users.length; i++){
              user.push(response.users[i].username)
              userFull.push(response.users[i]);
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/users?page_size=100", true);
      xhttp.send(); 
      };
      getUser();
    
      
    },
  data: function () {
        return {
            
            customerId: null,
            createdId: null,
            selectedDateFrom : null,
            selectedDateTo : null,
            showDialog: false,
            
            createdCustomer: customer,

            failAlert: false,
            successAlert: false,

            
            selectedCustomer: null,
            customerIdForm: null,
            costForm: null,
            priceForm: null,
            quantityForm: null,
            brandForm: null,
            positionForm: null,
            descriptionForm: null
       };
  },
  methods: {
     formSubmit(e){
       
          for (var i = 0; i < customerFull.length; i++){
          if (customerFull[i].customer_name == this.selectedCustomer){
            this.customerIdForm = customerFull[i].customer_id;
          }
        }
          axios.post('http://localhost:8081/api/bills',{
             product_name: this.productNameForm,
             group_id: this.groupIdForm,
             created_user_id: 1,//parseInt(localStorage.getItem("user_id")), /////// Fix this  
             cost: parseFloat(this.costForm),
             price: parseFloat(this.priceForm),
             quantity: parseInt(this.quantityForm),
             brand: this.brandForm,
             position: this.positionForm,
             description: this.descriptionForm
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
     }
  },
  components: {
    BillFilter,
    SimpleBill
  }
};
</script>
