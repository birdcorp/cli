import React, { useEffect } from 'react'
import {
  ChevronRight,
  Box,
  RefreshCw,
  Layers,
  X,
  ArrowBigLeft,
  Cpu,
  Globe,
  DollarSign,
  RectangleEllipsis,
  MousePointer,
  Bell,
  MessageSquare,
  Loader,
  Vibrate,
  TextSelect,
  Calendar,
  HelpCircle,
  Smile,
  Scan,
  TableProperties,
  SquareArrowUpRight,
  Phone,
  Copy,
  Clipboard,
  Truck,
  Play,
  Mic,
  Image,
  MapPin,
  Map,
  Crosshair,
  ArrowBigRight,
  PanelsTopLeft,
  ShoppingBasket,
  Eye,
  Database,
  KeyRound,
  ScanLine,
  Ticket,
  IterationCcw,
} from 'lucide-react'

import { LinkMiniProgram as Link } from '../Components'
import { useTab } from '../Context/Tabs'
import Tabbars from '../Components/Tabbars'

// Row component accepting icon component instead of string
function Row({
  icon: Icon,
  text,
  link,
}: {
  icon: React.ComponentType<any>
  text: string
  link: string
}) {
  return (
    <Link to={link}>
      <li className="w-full flex items-center justify-between mb-2 last:mb-0 border-b last:border-b-0 border-neutral-800 py-3">
        <div className="flex items-center">
          <span className="flex justify-center items-center mr-2 w-5 h-5 rounded-md">
            <Icon className="w-5 h-5" /> {/* Render Lucide icon here */}
          </span>
          <p>{text}</p>
        </div>
        <ChevronRight className="w-4 h-4 text-gray-400" />
      </li>
    </Link>
  )
}

export default function App() {
  const { setCurrentTab } = useTab()

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Developer',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <main className="pb-20 text-white">
      <div className="max-w-screen-sm mx-auto max-w-lg">
        <section className="mx-4 mt-4 mb-10">
          <div className="bg-neutral-950 h-56 mb-5 flex justify-center items-center rounded-md">
            <div className="flex flex-col">
              <p className="text-white">Mini Programs</p>
              <h3 className="text-purple-500 text-3xl">Developer Guide</h3>
            </div>
          </div>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">App</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Native bridge" icon={Layers} link="/app/bridge" />
            <Row text="App properties" icon={TableProperties} link="/app/id" />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Payments</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Pay" icon={DollarSign} link="/payment/checkout" />
            <Row
              text="Subscription"
              icon={IterationCcw}
              link="/payment/subscription"
            />
            <Row text="Coupon" icon={Ticket} link="/payment/coupon" />

            {/*<Row
              text="List orders"
              icon={RefreshCw}
              link="/payment/list-orders"
            />*/}
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Navigation</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row
              text="Navigation bar"
              icon={PanelsTopLeft}
              link="/navigation/navigation-bar"
            />
            <Row
              text="Navigate to"
              icon={ArrowBigRight}
              link="/navigation/navigate-to"
            />
            <Row
              text="Navigate back"
              icon={ArrowBigLeft}
              link="/navigation/navigate-back"
            />
            <Row text="Close" icon={X} link="/navigation/close" />
            <Row
              text="Navigate to program"
              icon={Cpu}
              link="/navigation/navigate-to-miniprogram"
            />
            <Row
              text="Open safari view"
              icon={Globe}
              link="/navigation/open-safari-view"
            />
            <Row
              text="Open browser"
              icon={Globe}
              link="/navigation/open-browser"
            />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Feedback</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row
              text="Focus"
              icon={MousePointer}
              link="/feedback/focus-element"
            />
            <Row text="Alert" icon={Bell} link="/feedback/alert" />
            <Row text="Toast" icon={MessageSquare} link="/feedback/toast" />
            <Row text="Alert" icon={Bell} link="/feedback/alert" />
            <Row text="Toast" icon={MessageSquare} link="/feedback/toast" />
            <Row text="Loading" icon={Loader} link="/feedback/loading" />
            <Row text="Vibrate" icon={Vibrate} link="/feedback/vibrate" />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Components</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Picker" icon={TextSelect} link="/feedback/picker" />
            <Row
              text="Date picker"
              icon={Calendar}
              link="/feedback/date-picker"
            />
            <Row text="Confirm" icon={HelpCircle} link="/feedback/confirm" />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Storage</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Storage" icon={Database} link="/storage/kvstore" />
            <Row text="Cart" icon={ShoppingBasket} link="/storage/cart" />
            <Row
              text="Recently viewed"
              icon={Eye}
              link="/storage/recent-viewed"
            />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Auth</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Face id" icon={Smile} link="/auth/faceid" />
            <Row
              text="Onetime code"
              icon={RectangleEllipsis}
              link="/auth/onetime-code"
            />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Device</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Key chain" icon={KeyRound} link="/device/keychain" />
            <Row text="Scan" icon={ScanLine} link="/device/scan" />
            <Row
              text="Make phone call"
              icon={Phone}
              link="/device/phone-call"
            />
            <Row
              text="Get clipboard"
              icon={Copy}
              link="/device/get-clipboard"
            />
            <Row
              text="Set clipboard"
              icon={Clipboard}
              link="/device/set-clipboard"
            />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Share</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Share" icon={SquareArrowUpRight} link="/share/link" />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Request user info</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Email" icon={MessageSquare} link="/user/email" />
            <Row text="Phone" icon={Phone} link="/user/phone" />
            <Row
              text="Shipping address"
              icon={Truck}
              link="/user/shipping-address"
            />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Media</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row text="Open video" icon={Play} link="/media/video" />
            <Row text="Open voice input" icon={Mic} link="/media/voice" />
            <Row text="Image picker" icon={Image} link="/media/image-picker" />
            <Row text="Upload files" icon={Box} link="/media/image-picker" />
          </ul>

          <h3 className="mb-2 ml-5 text-gray-400 mt-5">Location</h3>
          <ul className="text-white bg-neutral-950 px-5 rounded-md">
            <Row
              text="Get location"
              icon={Crosshair}
              link="/user-info/get-location"
            />
            <Row
              text="Open location"
              icon={MapPin}
              link="/user-info/open-location"
            />
            <Row
              text="Open map with directions"
              icon={Map}
              link="/user-info/open-map-with-directions"
            />
          </ul>
        </section>
      </div>

      <Tabbars onTabChange={setCurrentTab} />
    </main>
  )
}
