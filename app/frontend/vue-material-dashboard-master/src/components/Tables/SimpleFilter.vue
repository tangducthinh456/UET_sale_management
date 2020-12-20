<template>
  <div>
    <md-field >
      <label>Product Name</label>
      <md-input v-model="productName"  ></md-input>
    </md-field>
    
    <md-autocomplete  v-model="selectedGroup" :md-options="createdGroup">
      <label>Group</label>
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
  var group = [];
 var groupFull = [];
  var user = [];
  var userFull = [];
  
  //xhttp.open("GET", "ajax_info.txt", true);
  //xhttp.send();

  export default {
    name: 'TextFields',
    created: function(){
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
    data: () => ({
      initial: 'Initial Value',
      type: null,
      withLabel: null,
      inline: null,
      number: null,
      textarea: null,
      autogrow: null,
      disabled: null,
      productName: null,
      selectedGroup: null,
      selectedDateFrom: null,
      selectedDateTo: null,
      selectedCreated: null,
      createdId : null,
      groupId : null,
      createdGroup: group,
      createdUser: user,
    }),
    watch :{
      selectedGroup : function(){
        for (var i = 0; i < groupFull.length; i++){
          if (groupFull[i].group_name == this.selectedGroup){
            this.groupId = groupFull[i].group_id;
          }
        }
        this.$emit('update:groupId', this.groupId);
      },
      selectedCreated : function(){
        for (var i = 0; i < userFull.length; i++){
          if (userFull[i].username == this.selectedCreated){
            this.createdId = userFull[i].id;
          }
        }
        
        this.$emit('update:createdId', this.createdId);
      },
      productName : function(){
        this.$emit('update:productName', this.productName);
       
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