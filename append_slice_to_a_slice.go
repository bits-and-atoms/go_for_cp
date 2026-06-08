func pivotArray(v []int, pivot int) []int {

    lt := []int{}
    rt := []int{}
    mid := []int{}
    for _,val := range v{
        if val<pivot{
            lt = append(lt,val)
        }else if val>pivot{
            rt = append(rt,val)
        }else{
            mid = append(mid,val)
        }
    }
    lt = append(lt,mid...)
	// append expects individual element but mid is a slice
	// so ... is basically unpacking mid into single elements and one by one they are going
	// in cpp we do lt.insert(lt.end(),mid.begin(),mid.end())
    lt = append(lt,rt...)
    return lt
}