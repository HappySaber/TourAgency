import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include", // чтобы куки передавались (если JWT в куках)
      });

      const result = await response.json();

      if (response.ok) {
        alert(result.success || "Успешный вход");
        navigate("/consultation"); // переход на страницу консультаций
      } else {
        alert(result.error || "Ошибка входа");
      }
    } catch (error) {
      alert("Ошибка соединения с сервером");
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Вход</h2>
      <form onSubmit={handleSubmit}>
        <label htmlFor="email">Email:</label><br />
        <input
          type="email"
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        /><br /><br />

        <label htmlFor="password">Пароль:</label><br />
        <input
          type="password"
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        /><br /><br />

        <input type="submit" value="Войти" />
      </form>

      <p>Нет аккаунта? Попросите админа его вам создать :3</p>
    </div>
  );
}
