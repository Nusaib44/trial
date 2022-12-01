#include <stdio.h>
#include <limits.h>

int main()
{
   int i,j,k,n,count,first,second,a[10],temp;
   printf("enter the size");
   scanf("%d",&n);
   printf("enter the valuse");
   for(i=0;i<n;i++)
   {
       scanf("%d",&a[i]);
   }
   first=INT_MAX;
   second=INT_MAX;
   for(i=0;i<n;i++)
   {
       count=0;
       for(j=2;j<a[i];j++)
       {
           if(a[i]%j==0)
           {
               count++;
               break;
           }
       }
       if(count==0)
       {
           for(k=0;k<n;i++)
            {
            if(a[k]<first)
             {
                second=first;
                first=a[k];
             }
            else if(a[k]>first&&a[k]<second)
             {
                 second=a[k];
             }
            }
       }
   }
  
   if(second<7)
   {
        for(i=0;i<n;i++)
        {
            temp=0;
            for(j=0;j<n;j++)
             {
                if(a[i]==a[j]&&i!=j)
                {
                    temp++;
                    break;
                }
             }
             if(temp!=0)
             {
               if(a[j]>7)
               {
                   a[j]=0;
               }
             }
        }
        for(i=0;i<n;i++)
        {
            printf("%d",a[i]);
        }
   }
    return 0;
}