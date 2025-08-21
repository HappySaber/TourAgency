import React from "react";
import { Link } from "react-router-dom";
import "./Layout.css"; // стили можно вынести отдельно

export default function Layout({ children }) {
  return (
    <div>
      <header>
        <nav style={{ display: "flex", gap: "10px", alignItems: "center" }}>
          <Link to="/">Главная</Link>
          <Link to="/employee">Сотрудники</Link>
          <Link to="/position">Должности</Link>
          <Link to="/tours">Туры</Link>
          <Link to="/provider">Поставщики</Link>
          <Link to="/client">Клиенты</Link>
          <Link to="/consultation">Консультации</Link>
          <Link to="/service">Услуги</Link>
          <form
            action="/logout"
            method="POST"
            style={{ margin: 0 }}
          >
            <button type="submit">Выход</button>
          </form>
        </nav>
      </header>
      <main style={{ padding: "20px" }}>{children}</main>
    </div>
  );
}