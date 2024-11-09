import React, { createContext, useContext, useState, ReactNode } from 'react'

export enum Tab {
  Components = 'components',
  API = 'api',
  Docs = 'docs',
}

// Define the shape of the context data
interface TabContextType {
  currentTab: Tab
  setCurrentTab: React.Dispatch<React.SetStateAction<Tab>>
}

// Create the context with the correct type
const TabContext = createContext<TabContextType | undefined>(undefined)

// Custom hook for using the TabContext
export const useTab = (): TabContextType => {
  const context = useContext(TabContext)
  if (!context) {
    throw new Error('useTab must be used within a TabProvider')
  }
  return context
}

// Define props for the TabProvider
interface TabProviderProps {
  children: ReactNode
}

// Create the TabProvider to manage the state
export const TabProvider: React.FC<TabProviderProps> = ({ children }) => {
  const [currentTab, setCurrentTab] = useState<Tab>(Tab.Components)

  return (
    <TabContext.Provider value={{ currentTab, setCurrentTab }}>
      {children}
    </TabContext.Provider>
  )
}
