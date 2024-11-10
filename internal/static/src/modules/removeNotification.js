const removeDomElement = function (e) {
  e.target.remove();
};

export const initNotificationRemoval = function () {
  const notifications = document.querySelectorAll(
    'p[class*="notifications__"]',
  );

  if (!notifications.length) return null;

  for (let i = 0; i < notifications.length; i++) {
    notifications[i].onanimationend = removeDomElement;
  }
};
