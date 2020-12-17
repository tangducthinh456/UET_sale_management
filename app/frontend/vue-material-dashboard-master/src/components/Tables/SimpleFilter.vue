<template>
  <div>
    <md-field >
      <label>Product Name</label>
      <md-input v-model="productName"  @keyup="$emit('update:productName', productName);"></md-input>
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
 // var groupID = [];
  var user = [];
  //var userID = [];
  
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
            //console.log(response[0].group_name)
            for (var i = 0; i < response.length; i++){
              group.push(response[i].group_name)
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
            //console.log(this.responseText)
            for (var i = 0; i < response.users.length; i++){
              user.push(response.users[i].username)
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
    methods : {
        getFilter(){
          // for (var i = 0; i < )
          this.$emit("getFilter", productName);//, groupId, createdId, selectedDateFrom, selectedDateTo)
        }
    }
  }
</script>