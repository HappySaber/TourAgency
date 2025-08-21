import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";

export default function EditEmployee() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    email: "",
    firstName: "",
    lastName: "",
    middleName: "",
    address: "",
    phoneNumber: "",
    dateOfBirth: "",
    dateOfHiring: "",
    position: "",
  });

  const [positions, setPositions] = useState([]);

  // Загружаем данные сотрудника
  useEffect(() => {
    (async () => {
      try {
        const res = await fetch(`/api/employee/${id}`);
        const data = await res.json();
        if (res.ok) {
          setFormData({
            email: data.Employee.Email || "",
            firstName: data.Employee.FirstName || "",
            lastName: data.Employee.LastName || "",
            middleName: data.Employee.MiddleName || "",
            address: data.Employee.Address || "",
            phoneNumber: data.Employee.PhoneNumber || "",
            dateOfBirth: data.Employee.DateOfBirth
              ? data.Employee.DateOfBirth.split("T")[0]
              : "",
            dateOfHiring: data.Employee.DateOfHiring
              ? data.Employee.DateOfHiring.split("T")[0]
              : "",
            position: data.Employee.PositionID || "",
          });
          setPositions(data.Positions || []);
        } else {
          alert(data.error || "Ошибка загрузки данных сотрудника");
        }
      } catch (err) {
        alert("Ошибка: " + err.message);
      }
    })();
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const formatToTimestamp = (dateStr) => {
    return dateStr ? `${dateStr}T00:00:00Z` : null;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`/employee/edit/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...formData,
          dateOfBirth: formatToTimestamp(formData.dateOfBirth),
          dateOfHiring: formatToTimestamp(formData.dateOfHiring),
        }),
      });

      if (res.ok) {
        alert("Сотрудник обновлён");
        navigate("/employee");
      } else {
        const result = await res.json();
        alert(result.error || "Ошибка при обновлении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактирование сотрудника</h2>
      <form onSubmit={handleSubmit}>
        <label>Email:</label><br />
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Имя:</label><br />
        <input
          type="text"
          name="firstName"
          value={formData.firstName}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Фамилия:</label><br />
        <input
          type="text"
          name="lastName"
          value={formData.lastName}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Отчество:</label><br />
        <input
          type="text"
          name="middleName"
          value={formData.middleName}
          onChange={handleChange}
        /><br /><br />

        <label>Адрес:</label><br />
        <input
          type="text"
          name="address"
          value={formData.address}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Телефон:</label><br />
        <input
          type="tel"
          name="phoneNumber"
          pattern="[0-9]*"
          value={formData.phoneNumber}
          onChange={handleChange}
        /><br /><br />

        <label>Дата рождения:</label><br />
        <input
          type="date"
          name="dateOfBirth"
          value={formData.dateOfBirth}
          onChange={handleChange}
        /><br /><br />

        <label>Дата приёма на работу:</label><br />
        <input
          type="date"
          name="dateOfHiring"
          value={formData.dateOfHiring}
          onChange={handleChange}
        /><br /><br />

        <label>Должность:</label><br />
        <select
          name="position"
          value={formData.position}
          onChange={handleChange}
          required
        >
          <option value="">-- Выберите должность --</option>
          {positions.map((pos) => (
            <option key={pos.ID} value={pos.ID}>
              {pos.Name}
            </option>
          ))}
        </select><br /><br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
}
