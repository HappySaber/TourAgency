import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";

const EditService = ({ services }) => {
  const { id } = useParams(); // Предполагаем, что путь "/service/edit/:id"
  const navigate = useNavigate();

  const [name, setName] = useState("");
  const [price, setPrice] = useState("");

  useEffect(() => {
    // Инициализация формы на основе переданного списка услуг
    const service = services.find((s) => s.ID === parseInt(id, 10));
    if (service) {
      setName(service.Name);
      setPrice(service.Price);
    }
  }, [id, services]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const service = { id: parseInt(id, 10), name, price };

      const res = await fetch(`/service/edit/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(service),
      });

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Ошибка при обновлении");
      }

      alert("Услуга обновлена");
      navigate("/service");
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактировать услугу</h2>
      <form onSubmit={handleSubmit}>
        <label>Название:</label>
        <br />
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <br />
        <br />
        <label>Цена:</label>
        <br />
        <input
          type="text"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
          required
        />
        <br />
        <br />
        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
};

export default EditService;
