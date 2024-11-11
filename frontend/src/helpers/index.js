import {Notify} from "quasar";

export const helpers = () => {

  const message = (msg, type) => {
    const close = Notify.create({
      message: msg,
      type: type,
      position: "top-right",
      html: true,
      classes: "custom-notify",
      timeout: 5000, // Tiempo en milisegundos, 5000ms = 5 segundos
      actions: [
        {
          icon: "close", // Icono de cerrar
          color: "white", // Color del icono
          handler: () => {
            close()
          }
        }
      ],
      progress: true, // Barra de progreso visible
    });
  }

  return {
    message
  }
}
