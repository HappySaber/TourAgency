import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const CreateService = () => {
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const service = { name, price };

      const res = await fetch("/service/new", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(service),
      });

      const data = await res.json();

      if (!res.ok) throw new Error(data.error || data.message || "Неизвестная ошибка");

      alert(data.message || "Услуга успешно создана");
      navigate("/service"); // Перенаправление на список услуг
    } catch (err) {
      console.error("Ошибка при создании услуги:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавить услугу</h2>
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
        <input type="submit" value="Сохранить" />
      </form>
    </div>
  );
};

export default CreateService;
