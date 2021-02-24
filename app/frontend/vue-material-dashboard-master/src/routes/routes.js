import DashboardLayout from "@/pages/Layout/DashboardLayout.vue";

import Dashboard from "@/pages/Dashboard.vue";
import User from "@/pages/User.vue";
import TableList from "@/pages/TableList.vue";
import Typography from "@/pages/Typography.vue";
import Icons from "@/pages/Icons.vue";
import Maps from "@/pages/Maps.vue";
import Products from "@/pages/Products.vue";
import Notifications from "@/pages/Notifications.vue";
import UpgradeToPRO from "@/pages/UpgradeToPRO.vue";
import Bill from "@/pages/Bill.vue";
import Login from "@/pages/Login.vue"

const routes = [
  
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
    children: [
      
      {
        path: "dashboard",
        name: "Dashboard",
        component: Dashboard
      },
      {
        path: "user",
        name: "User",
        component: User
      },
      {
        path: "table",
        name: "Table List",
        component: TableList
      },
      {
        path: "product",
        name: "Products",
        component: Products
      },
      {
        path: "bill",
        name: "Bill",
        component: Bill
      },
      {
        path: "typography",
        name: "Typography",
        component: Typography
      },
      {
        path: "icons",
        name: "Icons",
        component: Icons
      },
      {
        path: "maps",
        name: "Maps",
        meta: {
          hideFooter: true
        },
        component: Maps
      },
      {
        path: "notifications",
        name: "Notifications",
        component: Notifications
      },
      {
        path: "upgrade",
        name: "Upgrade to PRO",
        component: UpgradeToPRO
      }
    ]
  },
  {
    path: "/login",
    name: "Login",
    component: Login  
  }
];

export default routes;
