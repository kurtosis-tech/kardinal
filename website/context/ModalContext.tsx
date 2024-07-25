"use client";

import { usePrevious } from "@uidotdev/usehooks";
import { usePathname } from "next/navigation";
import React, {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from "react";

import analytics from "@/lib/analytics";

interface ModalContextProps {
  isOpen: boolean;
  isNavOpen: boolean;
  toggleModal: () => void;
  toggleNav: () => void;
  // eslint-disable-next-line no-unused-vars
  setIsNavOpen: (newState: boolean) => void;
}

const ModalContext = createContext<ModalContextProps | undefined>(undefined);

export const ModalProvider = ({ children }: PropsWithChildren) => {
  const [isOpen, setIsOpen] = useState(false);
  const [isNavOpen, setIsNavOpen] = useState(false);
  const prevIsOpen = usePrevious(isOpen);
  const pathname = usePathname();

  const toggleModal = () => setIsOpen(!isOpen);
  const toggleNav = () => setIsNavOpen(!isNavOpen);

  // Track modal open and close events
  useEffect(() => {
    if (prevIsOpen === undefined) return;
    if (prevIsOpen === isOpen) return;
    if (isOpen && !prevIsOpen) {
      analytics.track("MODAL_OPENED");
    }
    if (!isOpen && prevIsOpen) {
      analytics.track("MODAL_CLOSED");
    }
  }, [isOpen, prevIsOpen]);

  // close the nav and modal when the route changes
  useEffect(() => {
    setIsNavOpen(false);
    setIsOpen(false);
  }, [pathname]);

  return (
    <ModalContext.Provider
      value={{
        isOpen,
        isNavOpen,
        toggleModal,
        toggleNav,
        setIsNavOpen,
      }}
    >
      {children}
    </ModalContext.Provider>
  );
};

export const useModal = (): ModalContextProps => {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error("useModal must be used within a ModalProvider");
  }
  return context;
};
