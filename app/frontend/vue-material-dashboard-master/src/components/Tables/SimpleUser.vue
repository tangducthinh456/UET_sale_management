<template>
  <div>
  
    <md-table v-model="users" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }" @click="showDialog = true; setForm(item)">
        <md-table-cell md-label="User Name">{{ item.username }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Role">{{ item.role }}</md-table-cell>
        <md-table-cell md-label="Phone Number">{{ item.phone_number }}</md-table-cell>
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
          <label>User Name</label>
         <md-input v-model="userNameForm"></md-input>   
     </md-field>

     
    <md-field >
          <label>Name</label>
         <md-input v-model="nameForm" ></md-input>
     </md-field>

     <md-field >
          <label>Password</label>
         <md-input v-model="password" type="password" ></md-input>
     </md-field>

   
          
         <md-autocomplete v-model="roleForm" :md-options="createdRole">
             <label>Role</label>
         </md-autocomplete>
    
    <md-field >
          <label>Phone number</label>
         <md-input v-model="phoneForm"></md-input>
     </md-field>

     <md-field >
          <label>Email</label>
         <md-input v-model="emailForm"></md-input>
     </md-field>
      </form>
      
      </div>
       </md-dialog-content>
    </md-dialog>

    
  </div>




  </div>
</template>

<script>
import axios from 'axios';

var users = [];
export default {
  name: "simple-user",
  created: function(){
       var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
            while(users.length > 0){ 
              users.pop();
            }
            //console.log("aaaaaaaaa");
            var response = JSON.parse(this.responseText)
            var varib = response.users;
            for (var i = 0; i < varib.length; i++){
              users.push(varib[i]);
              //prod.push({id: varib[i].product_id, name : varib[i].product_name, price : varib[i].price, quantity: varib[i].quantity})
            }
      }
      };
      xhttp.open("GET", "http://localhost:8081/api/users?page_size=1000", true);
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
     showDialog: false,
            
            createdUser: user,

            failAlert: false,
            successAlert: false,

            userNameForm: null,
            nameForm: null,
            roleForm: null,
            createdRole: role,
            phoneForm: null,
            emailForm: null,
            createdAtDisplay: null,
            password: null,
            userId: null
    };
  },
  
  methods: {
    setForm(item){
      this.userNameForm = item.username;
      this.nameForm = item.name;
      this.roleForm = item.role;
      this.phoneForm = item.phone_number;
      this.emailForm = item.email;
      this.userId = item.id;
      this.createdAtDisplay = item.created_at.replace(/T/g,' ').replace(/\+(.*)$/g,'');
      
    },
    updateProduct(){
       
      axios.put('http://localhost:8081/api/users/' + this.productIDForm,{
             username: this.userNameForm,
             name: this.nameForm,
             password: this.password,//parseInt(localStorage.getItem("user_id")), /////// Fix this  
             role: this.roleForm,
             email: this.emailForm,
             phone_number: this.phoneForm
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
          console.log(users);    
    },
    deleteProduct(){
      axios.delete('http://localhost:8081/api/users/' + this.userId)
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

  }
};
</script>
