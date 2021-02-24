<template>
  <div class="wrapper" :class="{ 'nav-open': $sidebar.showSidebar }">
    <notifications></notifications>
    <div id="message"></div>
    <side-bar
      :sidebar-item-color="sidebarBackground"
      :sidebar-background-image="sidebarBackgroundImage"
    >
      <mobile-menu slot="content"></mobile-menu>
      <sidebar-link to="/dashboard">
        <md-icon>dashboard</md-icon>
        <p>Dashboard</p>
      </sidebar-link>
      <sidebar-link to="/user">
        <md-icon>person</md-icon>
        <p>User Profile</p>
      </sidebar-link>
      <sidebar-link to="/product">
        <md-icon>shopping_cart</md-icon>
        <p>Product</p>
      </sidebar-link>
      <!-- <sidebar-link to="/import">
        <md-icon>library_books</md-icon>
        <p>Product</p>
      </sidebar-link>
      <sidebar-link to="/customer">
        <md-icon>shopping_cart</md-icon>
        <p>Product</p>
      </sidebar-link>
      <sidebar-link to="/provider">
        <md-icon>person</md-icon>
        <p>Product</p>
      </sidebar-link> -->
      <sidebar-link to="/bill">
        <md-icon>content_paste</md-icon>
        <p>Bill</p>
      </sidebar-link>
      <sidebar-link to="/typography">
        <md-icon>library_books</md-icon>
        <p>Typography</p>
      </sidebar-link>
      <sidebar-link to="/icons">
        <md-icon>bubble_chart</md-icon>
        <p>Icons</p>
      </sidebar-link>
      <sidebar-link to="/maps">
        <md-icon>location_on</md-icon>
        <p>Maps</p>
      </sidebar-link>
      <sidebar-link to="/notifications">
        <md-icon>notifications</md-icon>
        <p>Notifications</p>
      </sidebar-link>
      <sidebar-link to="/upgrade" class="active-pro">
        <md-icon>unarchive</md-icon>
        <p>Upgrade to PRO</p>
      </sidebar-link>
    </side-bar>

    <div class="main-panel">
      <top-navbar></top-navbar>

      <fixed-plugin
        :color.sync="sidebarBackground"
        :image.sync="sidebarBackgroundImage"
      >
      </fixed-plugin>

      <dashboard-content> </dashboard-content>

      <content-footer v-if="!$route.meta.hideFooter"></content-footer>
    </div>
  </div>
</template>

<script>


window.onload = function() {
  // Get a reference to the div on the page that will display the
  // message text.
  var messageEle = document.getElementById('message');

  // A function to process messages received by the window.
  function receiveMessage(e) {
    // Check to make sure that this message came from the correct domain.
    if (e.origin !== "http://localhost")
      return;

    // Update the div element to display the message.
    messageEle.innerHTML = "Message Received: " + e.data;
    localStorage.setItem("user_id", e.data.userID);
    localStorage.setItem("user_role", e.data.role);
  }

  // Setup an event listener that calls receiveMessage() when the window
  // receives a new MessageEvent.
  window.addEventListener('message', receiveMessage);
}



import TopNavbar from "./TopNavbar.vue";
import ContentFooter from "./ContentFooter.vue";
import DashboardContent from "./Content.vue";
import MobileMenu from "@/pages/Layout/MobileMenu.vue";
import FixedPlugin from "./Extra/FixedPlugin.vue";

export default {
  components: {
    TopNavbar,
    DashboardContent,
    ContentFooter,
    MobileMenu,
    FixedPlugin
  },
  data() {
    return {
      sidebarBackground: "green",
      sidebarBackgroundImage: require("@/assets/img/sidebar-2.jpg")
    };
  }
};
</script>
