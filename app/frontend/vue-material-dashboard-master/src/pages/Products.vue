<template>
  <div class="content">
    <div class="md-layout">
      <div
        class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
     
      <div>
    <md-dialog :md-active.sync="showDialog">
      <md-dialog-title>Product</md-dialog-title>
      <div style="margin: 30px">
      <form class="md-layout" method="POST" action="localhost:8081/api/products">
          
      <md-field>
          <label>Product Name</label>
         <md-input></md-input>
     </md-field>
    
    <md-field >
          <label>Cost ($)</label>
         <md-input></md-input>
     </md-field>

    <md-field >
          <label>Price ($)</label>
         <md-input></md-input>
     </md-field>

    <md-field >
          <label>Product Name</label>
         <md-input></md-input>
     </md-field>



      </form>
      </div>


      <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button class="md-primary" @click="showDialog = false">Save</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-button class="md-primary md-raised" @click="showDialog = true"><md-icon>library_add</md-icon>Create Product</md-button>
  </div>
      
        
      
        <md-card class="md-card-plain">
          <md-card-header data-background-color="green">
            <h4 class="title">Filter</h4>
            <p class="category">Here is a filter for this table</p>
          </md-card-header>
          <md-card-content>
            <simple-filter :productName.sync="productName" :createdId.sync="createdId" :groupId.sync="groupId" :selectedDateFrom.sync="selectedDateFrom" :selectedDateTo.sync="selectedDateTo"></simple-filter>
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
            <simple-product table-header-color="blue" :product-name="productName" :created-id="createdId" :group-id="groupId" :selected-date-from="selectedDateFrom" :selected-date-to="selectedDateTo"></simple-product>
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
            groupId: null,
            createdId: null,
            selectedDateFrom : null,
            selectedDateTo : null,
            showDialog: false,

      //       form: {
      //          firstName: null,
      //           lastName: null,
      //           gender: null,
      //            age: null,
      //             email: null,
      //        },
      // userSaved: false,
      // sending: false,
      // lastUser: null
       };
  },
  // validations: {
  //     form: {
  //       firstName: {
  //         required,
  //         minLength: minLength(3)
  //       },
  //       lastName: {
  //         required,
  //         minLength: minLength(3)
  //       },
  //       age: {
  //         required,
  //         maxLength: maxLength(3)
  //       },
  //       gender: {
  //         required
  //       },
  //       email: {
  //         required,
  //         email
  //       }
  //     }
  //   },
  components: {
    SimpleFilter,
    SimpleProduct
  }
};
</script>
