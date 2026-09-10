
import { createVNode, render, type Component } from 'vue'
import Popover from './index.vue'

type PopoverInstanceType = InstanceType<typeof Popover>

let popoverInstance: PopoverInstanceType | null = null


function getInstance(): PopoverInstanceType {
    if (!popoverInstance) {

        const container = document.createElement("div")
        document.body.append(container)
        const vNode = createVNode(Popover)

        render(vNode, container)

        popoverInstance = (vNode.component?.exposed as PopoverInstanceType) || null

    }
    return popoverInstance
}

export function usePopoverUtil(com: Component) {

    const instance = getInstance()


    const show = (el: HTMLElement, props?: Record<string, any>) => {
        instance.show(el, com, props)
    }


    return {
        show

    }


}