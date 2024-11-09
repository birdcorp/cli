import { Tab, useTab } from '../Context/Tabs'
import { Cog, FileCode, Book } from 'lucide-react'

interface Props {
  onTabChange: (tab: Tab) => void
}

export default function Tabbars({ onTabChange }: Props) {
  const { currentTab } = useTab()

  // Array of tab configurations
  const tabs = [
    { tab: Tab.Components, label: 'Components', Icon: Cog },
    { tab: Tab.API, label: 'API', Icon: FileCode },
    { tab: Tab.Docs, label: 'Docs', Icon: Book },
  ]

  return (
    <div className="fixed bottom-0 left-0 z-50 w-full h-20 bg-neutral-950 border-t border-neutral-800">
      <div className="grid h-full max-w-lg grid-cols-3 mx-auto font-medium">
        {tabs.map(({ tab, label, Icon }) => (
          <button
            key={label}
            type="button"
            onClick={() => onTabChange(tab)}
            className={`inline-flex flex-col items-center justify-center px-5 group ${
              currentTab === tab
                ? 'text-purple-600 dark:text-purple-500' // Active styles
                : 'text-gray-500 dark:text-gray-400 group-hover:text-purple-600 dark:group-hover:text-purple-500' // Default styles
            }`}
          >
            <Icon />
            <span
              className={`text-sm mb-4 ${
                currentTab === tab ? 'text-white' : ''
              }`}
            >
              {label}
            </span>
          </button>
        ))}
      </div>
    </div>
  )
}
