<template>
  <div class="content">
    <div class="md-layout">
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
     
      <div>
      <md-dialog-alert
      :md-active.sync="successAlert"
      md-content="Product has been created!"
      md-confirm-text="Ok" />

    <md-dialog-alert
      :md-active.sync="failAlert"
      md-title="Error!"
      md-content="Created Fail!" />
    <md-dialog :md-active.sync="showDialog">
      
      <md-dialog-title>User</md-dialog-title>
      
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
     
     <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button type="submit" class="md-primary" @click="showDialog = false;formSubmit()">Save</md-button>
      </md-dialog-actions>
      </form>
      
      </div>
       </md-dialog-content>
    </md-dialog>

    <md-button class="md-primary md-raised" @click="showDialog = true"><md-icon>library_add</md-icon>Create Product</md-button>
  </div>
      
        
      </div>
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100">
        <md-card>
          <md-card-header data-background-color="blue">
            <h4 class="title">User</h4>
            <p class="category">List user</p>
          </md-card-header>
          
          <md-card-content>
            <simple-user table-header-color="blue"></simple-user>
          </md-card-content>
        </md-card>
      </div>
    </div>
  </div>
</template>

<script >

import axios from 'axios';

  var user = [];
  var userFull = [];
  var role =["MANAGER", "SALE", "EMPLOYEE"];

import { SimpleUser } from "@/components";
//import Filter from '../components/Tables/Filter.vue';

export default {
  created: function(){

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
            createdAtForm: null,
            password: null,
       };
  },
  methods: {
     formSubmit(e){
       
          axios.post('http://localhost:8081/api/users ',{
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
     }
  },
  components: {
    SimpleUser
  }
};
</script>
