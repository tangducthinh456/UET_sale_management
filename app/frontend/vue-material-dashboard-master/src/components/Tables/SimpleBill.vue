<template>
  <div>
  
    <md-table v-model="bills" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }" @click="showDialog = true; setForm(item)">
        <md-table-cell md-label="Bill ID">{{ item.bill_id }}</md-table-cell>
        <md-table-cell md-label="Created At">{{ item.displayTime }}</md-table-cell>
        <md-table-cell md-label="Customer">{{ item.customer.customer_name }}</md-table-cell>
        <md-table-cell md-label="Total price ($)">{{ item.total }}</md-table-cell>
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
      
      <md-dialog-title>Bill</md-dialog-title>
      
      <md-dialog-content>
      
      <div style="margin: 30px">
      <form class="md-layout">
     
     <md-field >
          <label>Customer</label>
         <md-input v-model="selectedCustomer"></md-input>
     </md-field>

     <md-field >
          <label>Created User</label>
         <md-input v-model="createdUserForm"></md-input>
     </md-field>

     <md-field >
          <label>Created At</label>
         <md-input v-model="createdAtDisplay"></md-input>
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
    </md-dialog>

    
  </div>




  </div>
</template>

<script>

var bill = [];

import axios from 'axios';

var customer = [];
 var customerFull = [];

var prod = [];
export default {
  name: "simple-bill",
  created: function(){
       var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            while(bill.length > 0){ 
              bill.pop();
            }
            //console.log("aaaaaaaaa");
            var response = JSON.parse(this.responseText)
            var varib = response.bills
            for (var i = 0; i < varib.length; i++){
              varib[i].displayTime = varib[i].created_at.replace(/T/g,' ').replace(/\+(.*)$/g,'');
              varib[i].total = 0.0;
              const formatter = new Intl.NumberFormat('en-US', {
                       minimumFractionDigits: 2,      
                        maximumFractionDigits: 2,
                    });
              for (var j = 0; j < varib[i].details.length; j++){
                  varib[i].total += parseFloat(varib[i].details[j].product.cost) * varib[i].details[j].quantity;
              }
              varib[i].total = formatter.format(varib[i].total);
              bill.push(varib[i]);
              //prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/bills?page_size=100", true);
      xhttp.send(); 


      var getCustomer = function() {
          var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText)
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
      bills: bill,
      createdCustomer: customer,
      showDialog: false,

      failAlert: false,
      successAlert: false,

      
      
      selectedCustomer: null,
      customerIdForm: null,
      
      createdAtForm: null,
      createdUserForm: null,
      creatdUserIDForm: null,
      createdAtDisplay: null   
    };
  },
  props: ['createdId', 'selectedDateFrom', 'selectedDateTo', 'customerId'],
  watch: {
  customerId: 'reload',
  createdId: 'reload',
  selectedDateTo: 'reload',
  selectedDateFrom: 'reload',
  },
  methods: {
    setForm(item){
    //   this.productIDForm = item.product_id;
    //   this.productNameForm = item.product_name;

      this.selectedCustomer = item.customer.customer_name;
    //   this.groupIdForm = item.group.customer_id;
    //   this.costForm = item.cost.toString();
    //   this.priceForm = item.price.toString();
    //   this.quantityForm = item.quantity;
    //   this.brandForm = item.brand;
    //   this.positionForm = item.position;
    //   this.descriptionForm = item.description;
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
          //console.log(parseFloat(this.costForm) );    
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
            while(bill.length > 0){ 
              bill.pop();
            }
            //console.log("aaaaaaaaa");
            var response = JSON.parse(this.responseText)
            var varib = response.bills
            for (var i = 0; i < varib.length; i++){
              varib[i].displayTime = varib[i].created_at.replace(/T/g,' ').replace(/\+(.*)$/g,'');
              varib[i].total = 0.0;
              const formatter = new Intl.NumberFormat('en-US', {
                       minimumFractionDigits: 2,      
                        maximumFractionDigits: 2,
                    });
              for (var j = 0; j < varib[i].details.length; j++){
                  varib[i].total += parseFloat(varib[i].details[j].product.cost) * varib[i].details[j].quantity;
              }
              varib[i].total = formatter.format(varib[i].total);
              bill.push(varib[i]);
              //prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      var url = "http://localhost:8081/api/bills?page_size=2000&filter=";
      if (this.customerId) {
           url += "%3Bcustomer_id=" + this.customerId; 
      }
      if (this.createdId) {
           url += "%3Bcreated_user_id=" + this.createdId; 
      }
      if (this.selectedDateFrom) {
           //console.log(this.selectedDateFrom)
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
