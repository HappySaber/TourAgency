// src/App.jsx
import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Layout from "./components/Layout";
import Home from "./pages/Home";
import Login from "./pages/employee/login";

// Employee
import { EmployeeList, EditEmployee, CreateEmployee } from "./pages/employee";

// Positions
import { PositionList, EditPosition, CreatePosition } from "./pages/position";

// Providers
import { ProviderList, EditProvider, CreateProvider } from "./pages/provider";

// Services
import { ServiceList, EditService, CreateService } from "./pages/service";

// Tours
import { TourList, EditTour, CreateTour } from "./pages/tour";

// Consultations (если у тебя есть страницы консультаций)
import { ConsultationList, EditConsultation, CreateConsultation } from "./pages/consultation";

export default function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />

          {/* Employees */}
          <Route path="/employee" element={<EmployeeList />} />
          <Route path="/employee/employeeEdit/:id" element={<EditEmployee />} />
          <Route path="/employee/employeeCreate" element={<CreateEmployee />} />

          {/* Positions */}
          <Route path="/position" element={<PositionList />} />
          <Route path="/position/edit/:id" element={<EditPosition />} />
          <Route path="/position/new" element={<CreatePosition />} />

          {/* Providers */}
          <Route path="/provider" element={<ProviderList />} />
          <Route path="/provider/edit/:id" element={<EditProvider />} />
          <Route path="/provider/new" element={<CreateProvider />} />

          {/* Services */}
          <Route path="/service" element={<ServiceList />} />
          <Route path="/service/edit/:id" element={<EditService />} />
          <Route path="/service/new" element={<CreateService />} />

          {/* Tours */}
          <Route path="/tours" element={<TourList />} />
          <Route path="/tours/edit/:id" element={<EditTour />} />
          <Route path="/tours/new" element={<CreateTour />} />

          {/* Consultations */}
          <Route path="/consultation" element={<ConsultationList />} />
          <Route path="/consultation/edit/:id" element={<EditConsultation />} />
          <Route path="/consultation/new" element={<CreateConsultation />} />
        </Routes>
      </Layout>
    </Router>
  );
}
