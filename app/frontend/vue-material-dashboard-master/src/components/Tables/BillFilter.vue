<template>
  <div>
    
    <md-autocomplete  v-model="selectedCustomer" :md-options="createdCustomer">
      <label>Customer</label>
    </md-autocomplete>

    <md-autocomplete  v-model="selectedCreated" :md-options="createdUser">
      <label>Created User</label>
    </md-autocomplete>
    <md-datepicker v-model="selectedDateFrom" >
    <label>From</label>
    </md-datepicker>
    <md-datepicker v-model="selectedDateTo" >
    <label>To</label>
    </md-datepicker>
  </div>
</template>

<script>
  var customer = [];
 var customerFull = [];
  var user = [];
  var userFull = [];
  
  //xhttp.open("GET", "ajax_info.txt", true);
  //xhttp.send();

  export default {
    name: 'bill-filter',
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
            //console.log(response);
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
            while(user.length > 0){ 
              user.pop();
              userFull.pop();
            }
            for (var i = 0; i < response.users.length; i++){
              user.push(response.users[i].name)
              userFull.push(response.users[i]);
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/users?page_size=100", true);
      xhttp.send(); 
      };
      getUser();
    
      
    },
    data: () => ({
      initial: 'Initial Value',
      type: null,
      withLabel: null,
      inline: null,
      number: null,
      textarea: null,
      autogrow: null,
      disabled: null,
      
      selectedCustomer: null,
      selectedDateFrom: null,
      selectedDateTo: null,
      selectedCreated: null,
      createdId : null,
      customerId : null,
      createdCustomer: customer,
      createdUser: user
    }),
    watch :{
      selectedCustomer : function(){
        for (var i = 0; i < customerFull.length; i++){
          if (customerFull[i].customer_name == this.selectedCustomer){
            this.customerId = groupFull[i].customer_id;
          }
        }
        this.$emit('update:customerId', this.customerId);
      },
      selectedCreated : function(){
        for (var i = 0; i < userFull.length; i++){
          if (userFull[i].username == this.selectedCreated){
            this.createdId = userFull[i].id;
          }
        }
        
        this.$emit('update:createdId', this.createdId);
      },
      selectedDateFrom : function(){
        this.$emit('update:selectedDateFrom', this.selectedDateFrom);
       
      },
      selectedDateTo : function(){
        this.$emit('update:selectedDateTo', this.selectedDateTo);
       
      }
    }
  }
</script>