import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const EditTour = ({ tourData, providers }) => {
  const navigate = useNavigate();

  const [tour, setTour] = useState({
    id: tourData.ID,
    name: tourData.Name || "",
    rating: tourData.Rating || "",
    hotel: tourData.Hotel || "",
    nutrition: tourData.Nutrition || "",
    city: tourData.City || "",
    country: tourData.Country || "",
    provider: tourData.ProviderID || "",
    price: tourData.Price || "",
  });

  const [error, setError] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setTour((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    try {
      const res = await fetch(`/tours/edit/${tour.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(tour),
      });

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Ошибка при обновлении тура");
      }

      alert("Обновлено");
      navigate("/tours");
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h2>Редактировать тур</h2>
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

        <label>Цена:</label><br />
        <input type="text" name="price" value={tour.price} onChange={handleChange} /><br /><br />

        <label>Провайдер:</label><br />
        <select name="provider" value={tour.provider} onChange={handleChange} required>
          <option value="">-- Выберите провайдера --</option>
          {providers.map((p) => (
            <option key={p.ID} value={p.ID}>{p.Name}</option>
          ))}
        </select><br /><br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
};

export default EditTour;
