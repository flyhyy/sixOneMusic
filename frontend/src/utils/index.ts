

/* prettier-ignore */
/**
 * @Author: hyy
 * @Date: 2026-08-05 15:15:28
 * @Description: 
 */

/**
 * 获取列表中指定键值对应的索引
 * @param list 对象数组
 * @param pk 对象的键名 (Primary Key)
 * @param dfk 需要匹配的值 (Target Value)
 * @returns 匹配项的索引，未找到返回 -1
 */
export const getKeyIndex = <T extends Record<string, any>>(list: T[], pk: keyof T, targetVal: any): number => {
    return list.findIndex((item => item[pk] === targetVal))
}