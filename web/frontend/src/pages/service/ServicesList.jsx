import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const ServicesList = () => {
  const [services, setServices] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    // Получение списка услуг с сервера
    const fetchServices = async () => {
      try {
        const res = await fetch("/api/services"); // замените на ваш эндпоинт
        const data = await res.json();
        setServices(data.services || []);
      } catch (err) {
        console.error("Ошибка при загрузке услуг:", err);
        alert("Ошибка при загрузке услуг");
      }
    };

    fetchServices();
  }, []);

  const deleteService = async (id) => {
    if (!window.confirm("Удалить услугу?")) return;

    try {
      const res = await fetch(`/service/delete/${id}`, { method: "POST" });
      const data = await res.json();

      if (!res.ok) throw new Error(data.error || "Ошибка при удалении");

      // Обновляем список после удаления
      setServices((prev) => prev.filter((s) => s.ID !== id));
      alert("Удалено");
    } catch (err) {
      console.error("Ошибка при удалении услуги:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Список услуг</h2>
      <table border="1" cellPadding="5" cellSpacing="0">
        <thead>
          <tr>
            <th>Название</th>
            <th>Цена</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {services.length > 0 ? (
            services.map((service) => (
              <tr key={service.ID}>
                <td>{service.Name}</td>
                <td>{service.Price}</td>
                <td>
                  <button onClick={() => navigate(`/service/edit/${service.ID}`)}>
                    ✏️ Редактировать
                  </button>
                  <button onClick={() => deleteService(service.ID)}>🗑️ Удалить</button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="3">Услуги не найдены</td>
            </tr>
          )}
        </tbody>
      </table>
      <br />
      <button onClick={() => navigate("/service/new")}>➕ Добавить услугу</button>
    </div>
  );
};

export default ServicesList;
