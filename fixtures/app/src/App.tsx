import React, { Suspense } from 'react'
import LogRocket from 'logrocket'

import Tabbars from './Components/Tabbars'

import { useTab, Tab } from './Context/Tabs'

import { createBrowserRouter, RouterProvider } from 'react-router-dom'

const Home = React.lazy(() => import('./Screens/Home'))
const Alert = React.lazy(() => import('./Screens/Feedback/Alert'))
const Confirm = React.lazy(() => import('./Screens/Feedback/Confirm'))
const Loading = React.lazy(() => import('./Screens/Feedback/Loading'))
const Toast = React.lazy(() => import('./Screens/Feedback/Toast'))
const Vibrate = React.lazy(() => import('./Screens/Feedback/Vibrate'))
const Email = React.lazy(() => import('./Screens/UserInfo/Email'))
const Phone = React.lazy(() => import('./Screens/UserInfo/Phone'))
const Scan = React.lazy(() => import('./Screens/Device/Scan'))
const ShippingAddress = React.lazy(() =>
  import('./Screens/UserInfo/ShippingAddress'),
)
const NavigationBar = React.lazy(() =>
  import('./Screens/Navigation/NavigationBar'),
)
const NavigateTo = React.lazy(() => import('./Screens/Navigation/NavigateTo'))
const NavigateBack = React.lazy(() =>
  import('./Screens/Navigation/NavigateBack'),
)
const NavigateToMiniProgram = React.lazy(() =>
  import('./Screens/Navigation/NavigateToMiniProgram'),
)
const OpenBrowser = React.lazy(() => import('./Screens/Navigation/OpenBrowser'))
const GetClipboard = React.lazy(() => import('./Screens/Device/GetClipboard'))
const PhoneCall = React.lazy(() => import('./Screens/Device/PhoneCall'))
const SetClipboard = React.lazy(() => import('./Screens/Device/SetClipboard'))
const FaceID = React.lazy(() => import('./Screens/Auth/FaceID'))

/*
 * Payments
 */
const Checkout = React.lazy(() => import('./Screens/Payments/Checkout'))
const QuickPay = React.lazy(() => import('./Screens/Payments/QuickPay'))
const Subscription = React.lazy(() => import('./Screens/Payments/Subscription'))
const CollectShippingAddress = React.lazy(() =>
  import('./Screens/Payments/CollectShippingAddress'),
)
const Coupon = React.lazy(() => import('./Screens/Payments/Coupon'))
const Thankyou = React.lazy(() => import('./Screens/Payments/Thankyou'))
const ListOrders = React.lazy(() => import('./Screens/Payments/ListOrders'))

const Search = React.lazy(() => import('./Screens/Search'))
const ChangeLog = React.lazy(() => import('./Screens/Widgets/ChangeLog'))
const GetLocation = React.lazy(() => import('./Screens/Location/GetLocation'))
const OpenLocation = React.lazy(() => import('./Screens/Location/OpenLocation'))
const OpenMapWithDirections = React.lazy(() =>
  import('./Screens/Location/OpenMapWithDirections'),
)
const Models = React.lazy(() => import('./Screens/AR/Models'))
const Video = React.lazy(() => import('./Screens/Media/Video'))
const AppInfo = React.lazy(() => import('./Screens/App/Info'))
const Bridge = React.lazy(() => import('./Screens/App/Bridge'))
const Voice = React.lazy(() => import('./Screens/Media/Voice'))
const ImagePicker = React.lazy(() => import('./Screens/Media/ImagePicker'))
const KeyChain = React.lazy(() => import('./Screens/Auth/KeyChain'))
const Picker = React.lazy(() => import('./Screens/Feedback/Picker'))
const OpenSafariView = React.lazy(() =>
  import('./Screens/Navigation/OpenSafariView'),
)
const DatePicker = React.lazy(() => import('./Screens/Feedback/DatePicker'))
const Cart = React.lazy(() => import('./Screens/Storage/Cart'))
const FocusElement = React.lazy(() => import('./Screens/Feedback/FocusElement'))
const CloseMiniProgram = React.lazy(() =>
  import('./Screens/Navigation/CloseMiniProgram'),
)
const RecentViewed = React.lazy(() => import('./Screens/Storage/RecentViewed'))
const KVStore = React.lazy(() => import('./Screens/Storage/KVStore'))
const ShareLink = React.lazy(() => import('./Screens/Share/Link'))
const OnetimeCode = React.lazy(() => import('./Screens/Auth/OnetimeCode'))

LogRocket.init('3ouaid/miniprogram-developer')

const router = createBrowserRouter([
  { path: '/', element: <Home /> },
  { path: '/app/id', element: <AppInfo /> },
  { path: '/app/bridge', element: <Bridge /> },
  { path: '/feedback/focus-element', element: <FocusElement /> },
  { path: '/feedback/alert', element: <Alert /> },
  { path: '/feedback/picker', element: <Picker /> },
  { path: '/feedback/date-picker', element: <DatePicker /> },
  { path: '/feedback/confirm', element: <Confirm /> },
  { path: '/feedback/loading', element: <Loading /> },
  { path: '/feedback/toast', element: <Toast /> },
  { path: '/feedback/vibrate', element: <Vibrate /> },
  { path: '/user/email', element: <Email /> },
  { path: '/user/phone', element: <Phone /> },
  { path: '/user/shipping-address', element: <ShippingAddress /> },
  { path: '/navigation/navigation-bar', element: <NavigationBar /> },
  { path: '/navigation/close', element: <CloseMiniProgram /> },
  { path: '/navigation/navigate-to', element: <NavigateTo /> },
  { path: '/navigation/navigate-back', element: <NavigateBack /> },
  {
    path: '/navigation/navigate-to-miniprogram',
    element: <NavigateToMiniProgram />,
  },
  { path: '/navigation/open-browser', element: <OpenBrowser /> },
  { path: '/navigation/open-safari-view', element: <OpenSafariView /> },
  { path: '/device/get-clipboard', element: <GetClipboard /> },
  { path: '/device/scan', element: <Scan /> },
  { path: '/device/set-clipboard', element: <SetClipboard /> },
  { path: '/device/phone-call', element: <PhoneCall /> },
  { path: '/device/keychain', element: <KeyChain /> },
  { path: '/auth/faceid', element: <FaceID /> },
  { path: '/auth/onetime-code', element: <OnetimeCode /> },
  { path: '/storage/cart', element: <Cart /> },
  { path: '/storage/recent-viewed', element: <RecentViewed /> },
  { path: '/storage/kvstore', element: <KVStore /> },
  { path: '/payment/coupon', element: <Coupon /> },
  { path: '/payment/quickpay', element: <QuickPay /> },
  { path: '/payment/checkout', element: <Checkout /> },
  { path: '/payment/list-orders', element: <ListOrders /> },
  {
    path: '/payment/collect-shipping-address',
    element: <CollectShippingAddress />,
  },
  { path: '/payment/subscription', element: <Subscription /> },
  { path: '/payment/thankyou', element: <Thankyou /> },
  { path: '/search', element: <Search /> },
  { path: '/widgets/changelog', element: <ChangeLog /> },
  { path: '/user-info/get-location', element: <GetLocation /> },
  { path: '/user-info/open-location', element: <OpenLocation /> },
  {
    path: '/user-info/open-map-with-directions',
    element: <OpenMapWithDirections />,
  },
  { path: '/ar/models', element: <Models /> },
  { path: '/media/video', element: <Video /> },
  { path: '/media/voice', element: <Voice /> },
  { path: '/media/image-picker', element: <ImagePicker /> },
  { path: '/share/link', element: <ShareLink /> },
])

function App() {
  const { currentTab, setCurrentTab } = useTab()

  const renderContent = () => {
    switch (currentTab) {
      case Tab.Components:
        return <RouterProvider router={router} />
      case Tab.API:
        return (
          <div>
            <h3>Wallet Screen</h3>
            <Tabbars onTabChange={setCurrentTab} />
          </div>
        )
      case Tab.Docs:
        return (
          <div>
            <h3>Docs</h3>
            <Tabbars onTabChange={setCurrentTab} />
          </div>
        )
      default:
        return <h3>Home Screen</h3>
    }
  }

  return (
    <div className="container mx-auto min-h-screen">
      <Suspense fallback={<div>Loading...</div>}>{renderContent()}</Suspense>
    </div>
  )
}

export default App
