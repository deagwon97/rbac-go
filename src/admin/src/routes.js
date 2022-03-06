// import { Navigate } from 'react-router-dom';
import MainPage from "pages/MainPage.js";
// import ViewContentPage from 'pages/post/ViewContentPage';
// import EditContentPage from 'pages/post/EditContentPage';
// import CreateContentPage from 'pages/post/CreateContentPage';
// import CreateRoomPage from 'pages/chat/CreateRoomPage';
// import RoomPage from 'pages/chat/RoomPage';
// import SignUpPage from 'pages/accounts/SignUpPage';
// import KakaoLoginPage from 'pages/accounts/KakaoLoginPage';

const routes = [
  {
    path: "",
    element: <MainPage />,
  },
  // element: <Navigate to="/list" />,

  // {
  // 	path: '/content',
  // 	element: <ViewContentPage />,
  // 	// children: [{ path: 'edit', element: <EditContentPage /> }],
  // },
  // {
  // 	path: '/edit',
  // 	element: <EditContentPage />,
  // },

  // {
  // 	path: '/create',
  // 	element: <CreateContentPage />,
  // },
  // {
  // 	path: '/chat',
  // 	element: <CreateRoomPage />,
  // },
  // {
  // 	path: '/chat/room/',
  // 	element: <RoomPage />,
  // },
  // {
  // 	path: '/signup/',
  // 	element: <SignUpPage />,
  // },
  // {
  // 	path: '/accounts/kakao/callback/',
  // 	element: <KakaoLoginPage />,
  // },

  // 	children: [{ path: 'input', element: <AddViewPage /> }],
  // },
  // {
  // 	path: 'add',
  // 	element: <AddViewPage />,
  // },
  // {
  // 	path: 'input',
  // 	element: <InputViewPage />,
  // },
  // {
  // 	path: 'blog',
  // 	element: <DashboardLayout />,
  // 	children: [
  // 		{ path: 'view', element: <ContentViewPage /> },
  // 		{ path: 'edit', element: <ContentEditPage /> },
  // 		{ path: 'list', element: <ContentListPage /> },
  // 	],
  // },
  // {
  // 	path: '/',
  // 	element: <DashboardLayout />,
  // 	children: [
  // 		{ path: '/', element: <Navigate to="/app/main" /> },
  // 		// { path: '*', element: <Navigate to="/404" /> },
  // 	],
  // },
];

export default routes;
