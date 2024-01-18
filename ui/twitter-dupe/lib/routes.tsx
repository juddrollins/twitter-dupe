const ROUTES = [
  {
    page: "/",
    regex: "^/(?:/)?$",
    routeKeys: {},
    namedRegex: "^/(?:/)?$",
  },
  {
    page: "/_not-found",
    regex: "^/_not\\-found(?:/)?$",
    routeKeys: {},
    namedRegex: "^/_not\\-found(?:/)?$",
  },
  {
    page: "/debug",
    regex: "^/debug(?:/)?$",
    routeKeys: {},
    namedRegex: "^/debug(?:/)?$",
  },
  {
    page: "/favicon.ico",
    regex: "^/favicon\\.ico(?:/)?$",
    routeKeys: {},
    namedRegex: "^/favicon\\.ico(?:/)?$",
  },
  {
    page: "/home",
    regex: "^/home(?:/)?$",
    routeKeys: {},
    namedRegex: "^/home(?:/)?$",
  },
  {
    page: "/login",
    regex: "^/login(?:/)?$",
    routeKeys: {},
    namedRegex: "^/login(?:/)?$",
  },
];

const isRouteValid = (route: string) => {
  return ROUTES.some((r) => r.page === route);
};

export default isRouteValid;
