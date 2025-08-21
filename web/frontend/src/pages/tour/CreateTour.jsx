import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const CreateTour = ({ providers }) => {
  const navigate = useNavigate();

  const [tour, setTour] = useState({
    name: "",
    rating: "",
    hotel: "",
    nutrition: "",
    city: "",
    country: "",
    provider: "",
    price: "",
  });

  const [error, setError] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setTour((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    // базовая валидация
    if (!tour.name || !tour.provider) {
      setError("Поля 'Название' и 'Провайдер' обязательны");
      return;
    }

    try {
      const res = await fetch("/tours/new", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(tour),
      });

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Ошибка при создании тура");
      }

      alert("Тур создан");
      navigate("/tours");
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h2>Добавить тур</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <label>Название:</label><br />
        <input type="text" name="name" value={tour.name} onChange={handleChange} required /><br /><br />

        <label>Рейтинг:</label><br />
        <input type="text" name="rating" value={tour.rating} onChange={handleChange} /><br /><br />

        <label>Отель:</label><br />
        <input type="text" name="hotel" value={tour.hotel} onChange={handleChange} /><br /><br />

        <label>Питание:</label><br />
        <input type="text" name="nutrition" value={tour.nutrition} onChange={handleChange} /><br /><br />

        <label>Город:</label><br />
        <input type="text" name="city" value={tour.city} onChange={handleChange} /><br /><br />

        <label>Страна:</label><br />
        <input type="text" name="country" value={tour.country} onChange={handleChange} /><br /><br />

        <label>Цена за единицу:</label><br />
        <input type="text" name="price" value={tour.price} onChange={handleChange} /><br /><br />

        <label>Провайдер:</label><br />
        <select name="provider" value={tour.provider} onChange={handleChange} required>
          <option value="">-- Выберите провайдера --</option>
          {providers.map((p) => (
            <option key={p.ID} value={p.ID}>{p.Name}</option>
          ))}
        </select><br /><br />

        <button type="submit">Создать</button>
      </form>
    </div>
  );
};

export default CreateTour;
