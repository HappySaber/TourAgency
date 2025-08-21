import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const CreateProvider = () => {
  const [form, setForm] = useState({
    name: "",
    addressto: "",
    address: "",
    email: "",
    phonenumber: "",
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("/provider/new", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(form),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || data.message || "Неизвестная ошибка");
      }

      alert(data.message || "Поставщик успешно создан");
      navigate("/provider"); // переход на список провайдеров
    } catch (err) {
      console.error("Ошибка при создании поставщика:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавить провайдера</h2>
      <form onSubmit={handleSubmit}>
        <label>Название:</label><br />
        <input type="text" name="name" value={form.name} onChange={handleChange} required /><br /><br />

        <label>Адрес назначения:</label><br />
        <input type="text" name="addressto" value={form.addressto} onChange={handleChange} required /><br /><br />

        <label>Адрес:</label><br />
        <input type="text" name="address" value={form.address} onChange={handleChange} required /><br /><br />

        <label>Email:</label><br />
        <input type="email" name="email" value={form.email} onChange={handleChange} required /><br /><br />

        <label>Телефон:</label><br />
        <input type="text" name="phonenumber" value={form.phonenumber} onChange={handleChange} required /><br /><br />

        <button type="submit">Сохранить</button>
      </form>
    </div>
  );
};

export default CreateProvider;
